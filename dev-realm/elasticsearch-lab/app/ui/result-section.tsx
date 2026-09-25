import { SearchResult } from "../types"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Badge } from "@/components/ui/badge"
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Empty,
  EmptyDescription,
  EmptyMedia,
  EmptyTitle,
} from "@/components/ui/empty"
import ProductList from "./product-list"
import { HugeiconsIcon } from "@hugeicons/react"
import { Clock01Icon } from "@hugeicons/core-free-icons"

export default function ResultSection({
  title,
  searchResult,
}: {
  title: string
  searchResult?: SearchResult
}) {
  return (
    <Card className="h-full">
      <CardHeader>
        <CardTitle>{title}</CardTitle>
      </CardHeader>

      <CardContent className="flex-1">
        {searchResult ? (
          searchResult.error ? (
            <Alert variant="destructive">
              <AlertDescription>{searchResult.error}</AlertDescription>
            </Alert>
          ) : (
            <ProductList products={searchResult.products} />
          )
        ) : (
          <Empty className="min-h-40 border">
            <EmptyMedia variant="icon">
              <HugeiconsIcon icon={Clock01Icon} strokeWidth={1.5} />
            </EmptyMedia>
            <EmptyTitle>Waiting for search.</EmptyTitle>
            <EmptyDescription>
              Write something in the search bar and hit Enter key.
            </EmptyDescription>
          </Empty>
        )}
      </CardContent>

      {searchResult && !searchResult.error && (
        <CardFooter className="justify-between gap-2 border-t">
          <Badge variant="secondary">Latency: {searchResult.latency}ms</Badge>
          <Badge variant="outline">Found: {searchResult.count}</Badge>
        </CardFooter>
      )}
    </Card>
  )
}
