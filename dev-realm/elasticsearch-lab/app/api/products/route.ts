import { Client } from "@elastic/elasticsearch"
import { HttpConnection } from "@elastic/transport"
import { Product, SearchResult } from "@/app/types"
import postgres from "postgres"
import { NextRequest } from "next/server"

const ELASTICSEARCH_INDEX_NAME = "products"

const sql = postgres(process.env.POSTGRES_URL!, { max: 1 })
const esClient = new Client({
  node: process.env.ELASTICSEARCH_NODE_URL!,
  Connection: HttpConnection,
})

const roundOff = (ms: number) => Number.parseFloat(ms.toFixed(2))

type ProductRow = Product & { total_count: number }

async function searchPostgres(term: string): Promise<SearchResult> {
  const source = "PostgreSQL (ILIKE)"
  const pattern = `%${term}%`
  const start = performance.now()

  try {
    const rows = await sql<ProductRow[]>`
      SELECT id, internal_id AS "internalId", name, description, brand, category,
             price, currency, stock, ean, color, size, availability,
             COUNT(*) OVER() AS total_count
      FROM products
      WHERE name ILIKE ${pattern}
         OR brand ILIKE ${pattern}
         OR category ILIKE ${pattern}
         OR description ILIKE ${pattern}
         OR color ILIKE ${pattern}
         OR size ILIKE ${pattern}
      LIMIT 50
    `

    const total = rows[0]?.total_count ?? 0
    const products: Product[] = rows.map((row) => {
      const { total_count, ...product } = row
      void total_count
      return product
    })

    return {
      source,
      products,
      latency: roundOff(performance.now() - start),
      count: total,
      error: null,
    }
  } catch (error: unknown) {
    return {
      source,
      products: [],
      latency: roundOff(performance.now() - start),
      count: 0,
      error:
        error instanceof Error
          ? error.message
          : "Failed to fetch from PostgreSQL",
    }
  }
}

async function searchElastic(term: string): Promise<SearchResult> {
  const source = "Elasticsearch"
  const start = performance.now()

  try {
    const result = await esClient.search({
      index: ELASTICSEARCH_INDEX_NAME,
      size: 50,
      track_total_hits: true,
      query: {
        bool: {
          should: [
            {
              wildcard: {
                "name.raw": { value: `*${term}*`, case_insensitive: true },
              },
            },
            {
              wildcard: {
                "brand.raw": { value: `*${term}*`, case_insensitive: true },
              },
            },
            {
              wildcard: {
                "category.raw": { value: `*${term}*`, case_insensitive: true },
              },
            },
            {
              wildcard: {
                "description.raw": {
                  value: `*${term}*`,
                  case_insensitive: true,
                },
              },
            },
            {
              wildcard: {
                "color.raw": { value: `*${term}*`, case_insensitive: true },
              },
            },
            {
              wildcard: {
                "size.raw": { value: `*${term}*`, case_insensitive: true },
              },
            },
          ],
          minimum_should_match: 1,
        },
      },
    })

    const products: Product[] = result.hits.hits.map((hit) => ({
      ...(hit._source as Product),
      id: Number(hit._id),
    }))

    const total =
      typeof result.hits.total === "number"
        ? result.hits.total
        : result.hits.total?.value

    return {
      source,
      products,
      latency: roundOff(performance.now() - start),
      count: total as number,
      error: null,
    }
  } catch (error: unknown) {
    return {
      source,
      products: [],
      latency: roundOff(performance.now() - start),
      count: 0,
      error:
        error instanceof Error
          ? error.message
          : "Failed to fetch from Elasticsearch",
    }
  }
}

export async function GET(request: NextRequest) {
  const { searchParams } = request.nextUrl

  const query = searchParams.get("q")
  if (typeof query !== "string" || query.trim() === "") {
    return new Response(
      JSON.stringify({ error: "Search query is required." }),
      {
        status: 400,
        headers: {
          "Content-Type": "application/json",
        },
      }
    )
  }

  // Escape each query syntax independently so both searches match literal input.
  const pgTerm = query.replace(/[\\%_]/g, "\\$&")
  const esTerm = query.replace(/[\\*?]/g, "\\$&")

  const encoder = new TextEncoder()
  const searches = [
    [searchPostgres, pgTerm] as const,
    [searchElastic, esTerm] as const,
  ]

  const stream = new ReadableStream({
    async start(controller) {
      try {
        await Promise.all(
          searches.map(async ([search, term]) => {
            const result = await search(term)
            controller.enqueue(encoder.encode(JSON.stringify(result) + "\n"))
          })
        )
        controller.close()
      } catch (error) {
        controller.error(error)
      }
    },
  })

  return new Response(stream, {
    headers: {
      "Content-Type": "application/x-ndjson; charset=utf-8",
      "Cache-Control": "no-cache, no-transform",
    },
  })
}
