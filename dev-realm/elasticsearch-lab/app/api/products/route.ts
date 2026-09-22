import { Client } from "@elastic/elasticsearch"
import { HttpConnection } from "@elastic/transport"
import { Product, SearchResult } from "@/app/types"
import postgres from "postgres"

const ELASTICSEARCH_INDEX_NAME = "products"

const sql = postgres(process.env.POSTGRES_URL!, { max: 1 })
const esClient = new Client({
  node: process.env.ELASTICSEARCH_NODE_URL!,
  Connection: HttpConnection,
})

const roundOff = (ms: number) => Number.parseFloat(ms.toFixed(2))

async function searchPostgres(term: string): Promise<SearchResult> {
  const source = "PostgreSQL (ILIKE)"
  const pattern = `%${term}%`
  const start = performance.now()

  try {
    const rows = await sql<Product[]>`
      SELECT id, name, description, brand, category,
             price, currency, stock, ean, color, size, availability
      FROM products
      WHERE name ILIKE ${pattern}
         OR brand ILIKE ${pattern}
         OR category ILIKE ${pattern}
         OR description ILIKE ${pattern}
         OR color ILIKE ${pattern}
         OR size ILIKE ${pattern}
      LIMIT 50
    `

    return {
      source,
      products: [...rows],
      latency: roundOff(performance.now() - start),
      count: rows.length,
      error: null,
    }
  } catch (error: any) {
    return {
      source,
      products: [],
      latency: roundOff(performance.now() - start),
      count: 0,
      error: error?.message || "Failed to fetch from PostgreSQL",
    }
  }
}

async function searchElastic(term: string): Promise<SearchResult> {
  const source = "Elasticsearch"
  const start = performance.now()

  try {
    const result = await esClient.search({
      index: ELASTICSEARCH_INDEX_NAME,
      query: {
        query_string: {
          query: `*${term.toLowerCase()}*`,
          fields: ["name", "brand", "category", "description", "color", "size"],
          default_operator: "OR",
          analyze_wildcard: true,
        },
      },
    })

    const products: Product[] = result.hits.hits.map((hit) => ({
      ...(hit._source as Product),
      id: Number(hit._id),
    }))

    return {
      source,
      products,
      latency: roundOff(performance.now() - start),
      count: products.length,
      error: null,
    }
  } catch (error: any) {
    return {
      source,
      products: [],
      latency: roundOff(performance.now() - start),
      count: 0,
      error: error?.message || "Failed to fetch from Elasticsearch",
    }
  }
}

export async function GET() {
  const encoder = new TextEncoder()
  const searches = [searchPostgres, searchElastic]

  const stream = new ReadableStream({
    async start(controller) {
      try {
        await Promise.all(
          searches.map(async (search) => {
            const result = await search("ball")
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
