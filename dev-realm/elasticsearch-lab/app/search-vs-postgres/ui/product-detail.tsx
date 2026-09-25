import { Product } from "@/app/types"

export default function ProductDetail({ product }: { product: Product }) {
  return (
    <div className="space-y-1">
      <p className="text-sm font-medium">{product.name}</p>
      <p className="line-clamp-2 text-xs text-muted-foreground">
        {product.description}
      </p>

      <div className="flex items-center justify-between gap-2 text-xs text-muted-foreground">
        <p>
          {product.brand} - {product.category}
        </p>
        <p className="shrink-0 font-medium text-foreground">
          {product.currency} {product.price}
        </p>
      </div>
    </div>
  )
}
