import { Product } from "../types"
import { ScrollArea } from "@/components/ui/scroll-area"
import ProductDetail from "./product-detail"

export default function ProductList({ products }: { products: Product[] }) {
  return (
    <ScrollArea className="h-96 rounded-lg border">
      <div className="flex flex-col gap-1.5 p-1.5">
        {products.map((product: Product) => (
          <div key={product.id} className="rounded-md bg-muted px-3 py-2.5">
            <ProductDetail product={product} />
          </div>
        ))}
      </div>
    </ScrollArea>
  )
}
