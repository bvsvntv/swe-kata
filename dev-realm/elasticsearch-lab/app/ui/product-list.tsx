import { Product } from "../types"
import ProductDetail from "./product-detail"

export default function ProductList({ products }: { products: Product[] }) {
  return (
    <div className="mt-4">
      <div className="mx-auto w-full overflow-hidden rounded-2xl border">
        <div className="h-96 scroll-fade scrollbar-none overflow-y-auto">
          <div className="flex flex-col gap-1.5 p-1.5">
            {products.map((product: Product) => (
              <div
                key={product.id}
                className="rounded-lg bg-muted px-3 py-2.5 text-sm"
              >
                <ProductDetail product={product} />
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  )
}
