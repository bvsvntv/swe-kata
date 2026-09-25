import { ScrollArea } from "@/components/ui/scroll-area"
import ProductDetail from "./product-detail"
import { Product } from "@/app/types"

export default function ProductList({ products }: { products: Product[] }) {
  return (
    <ScrollArea className="h-96 rounded-lg">
      <div className="flex flex-col gap-1.5">
        {products.length > 0 ? (
          products.map((product: Product) => (
            <div key={product.id} className="rounded-md bg-muted px-3 py-2.5">
              <ProductDetail product={product} />
            </div>
          ))
        ) : (
          <div>
            <p className="text-xs/relaxed text-muted-foreground">
              No products found.
            </p>
          </div>
        )}
      </div>
    </ScrollArea>
  )
}
