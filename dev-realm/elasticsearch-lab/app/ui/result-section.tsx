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
      <h3 className="font-semibold">{title}</h3>

      <div className="">
        {searchResult ? (
          <div>
            {searchResult.error ? (
              <p className="text-sm text-red-500">{searchResult.error}</p>
            ) : (
              <ProductList products={searchResult.products} />
            )}

            <p>latency: {searchResult.latency}ms</p>
            <p>count: {searchResult.count}</p>
          </div>
        ) : (
          <p className="text-gray-500">Waiting for search.</p>
        )}
      </div>
    </section>
  )
}
