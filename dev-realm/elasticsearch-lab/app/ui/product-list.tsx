import { Product } from "../types"
import ProductDetail from "./product-detail"

export default function ProductList({ products }: { products: Product[] }) {
  return (
    <div className="mt-2">
      {products.map((product: Product) => (
        <div key={product.id} className="mt-2 rounded border p-0.5">
          <ProductDetail product={product} />
        </div>
      ))}
    </div>
  )
}
