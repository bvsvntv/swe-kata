"use client"

import { useState } from "react"
import { SearchResult } from "./types"
import ProductList from "./ui/product-list"

export default function Page() {
  const [searchResults, setSearchResults] = useState<SearchResult[]>([])

  const pgResults = searchResults.find(
    (product: SearchResult) => product.source === "PostgreSQL (ILIKE)"
  )

  const esResults = searchResults.find(
    (product: SearchResult) => product.source === "Elasticsearch"
  )

  async function handleSearch() {
    const response = await fetch("/api/products")

    const reader = response?.body!.getReader()
    const decoder = new TextDecoder()

    let buffer = ""
    while (true) {
      const { value, done } = await reader.read()
      if (done) break

      buffer += decoder.decode(value, { stream: true })
      const lines = buffer.split("\n")
      buffer = lines.pop() ?? ""

      for (const line of lines) {
        if (!line.trim()) continue

        const result = JSON.parse(line) as SearchResult
        setSearchResults((prev) => [...prev, result])

        console.log(result)
      }

      // Flush TextDecoder
      buffer += decoder.decode()

      if (buffer.trim()) {
        const result = JSON.parse(buffer)
        console.log(result)
      }
    }
  }

  return (
    <div className="m-4 mx-auto max-w-5xl">
      <button onClick={handleSearch} className="rounded bg-gray-700 p-2">
        search products
      </button>

      <div className="mt-6 flex justify-between">
        <section>
          <h3 className="font-semibold">PostgreSQL Search</h3>

          <div className="">
            {pgResults ? (
              <div>
                {pgResults.error ? (
                  <p className="text-sm text-red-500">{pgResults.error}</p>
                ) : (
                  <ProductList products={pgResults.products} />
                )}

                <p>latency: {pgResults.latency}ms</p>
                <p>count: {pgResults.count}</p>
              </div>
            ) : (
              <p className="text-gray-500">Waiting for search.</p>
            )}
          </div>
        </section>

        <section>
          <h3 className="font-semibold">Elasticsearch</h3>

          <div className="">
            {esResults ? (
              <div>
                {esResults.error ? (
                  <p className="text-sm text-red-500">{esResults.error}</p>
                ) : (
                  <ProductList products={esResults.products} />
                )}

                <p>latency: {esResults.latency}ms</p>
                <p>count: {esResults.count}</p>
              </div>
            ) : (
              <p className="text-gray-500">Waiting for search.</p>
            )}
          </div>
        </section>
      </div>
    </div>
  )
}
