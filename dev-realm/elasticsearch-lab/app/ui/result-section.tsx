import { SearchResult } from "../types"
import ProductList from "./product-list"

export default function ResultSection({
  title,
  searchResult,
}: {
  title: string
  searchResult: SearchResult
}) {
  return (
    <section>
      <h3 className="mt-4 font-semibold">{title}</h3>

      <div className="mt-2">
        {searchResult ? (
          <div className="mt-2">
            {searchResult.error ? (
              <p className="text-sm text-red-500">{searchResult.error}</p>
            ) : (
              <ProductList products={searchResult.products} />
            )}

            <div className="mt-2 flex justify-between">
              <p className="text-xs">Latency: {searchResult.latency}ms</p>
              <p className="text-xs">Found: {searchResult.count}</p>
            </div>
          </div>
        ) : (
          <p className="text-gray-500">Waiting for search.</p>
        )}
      </div>
    </section>
  )
}
