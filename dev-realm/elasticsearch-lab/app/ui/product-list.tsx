import { Product } from "../types"

export default function ProductList({ products }: { products: Product[] }) {
  return (
    <div className="mt-2">
      {products.map((product: Product) => (
        <div key={product.id} className="mt-2 rounded border p-0.5">
          <p className="text-sm font-normal">{product.name}</p>

          <div className="flex justify-between">
            <p className="text-xs">
              {product.brand} - {product.category}
            </p>
            <p className="text-xs">
              {product.currency} - {product.price}
            </p>
          </div>
        </div>
      ))}
    </div>
  )
}
