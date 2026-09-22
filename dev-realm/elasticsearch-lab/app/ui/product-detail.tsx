import { Product } from "../types"

export default function ProductDetail({ product }: { product: Product }) {
  return (
    <div>
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
  )
}
