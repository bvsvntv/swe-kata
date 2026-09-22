import { Product, SearchResult } from "@/app/types"
import postgres from "postgres"

const roundOff = (ms: number) => Number.parseFloat(ms.toFixed(2))

const sql = postgres(process.env.POSTGRES_URL!, { max: 1 })

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
    // TODO: call Elasticsearch here and map hits to Product-shaped objects
    throw new Error("Not implemented")
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
      } finally {
        controller.close()
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
