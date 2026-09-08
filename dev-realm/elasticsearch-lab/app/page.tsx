import { HugeiconsIcon } from "@hugeicons/react"
import { Button } from "@/components/ui/button"
import {
  Empty,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
  EmptyDescription,
  EmptyContent,
} from "@/components/ui/empty"
import { Field } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { SearchAlertIcon } from "@hugeicons/core-free-icons"

export default function Page() {
  return (
    <div className="flex min-h-svh p-6">
      <div className="mx-auto flex max-w-7xl min-w-md flex-col gap-4 text-sm leading-loose">
        <div>
          <h1 className="font-medium">
            Text search in Database V/S Elastisearch
          </h1>

          <Field className="mt-4" orientation="horizontal">
            <Input type="search" placeholder="Search..." />
            <Button>Search</Button>
          </Field>
        </div>

        <div>
          <Empty>
            <EmptyHeader>
              <EmptyMedia variant="icon">
                <HugeiconsIcon
                  icon={SearchAlertIcon}
                  size={24}
                  color="currentColor"
                  strokeWidth={1.5}
                />
              </EmptyMedia>
              <EmptyTitle>Waiting for search</EmptyTitle>
              <EmptyDescription>
                Enter a search term and click "Search" to see results.
              </EmptyDescription>
            </EmptyHeader>
          </Empty>
        </div>
      </div>
    </div>
  )
}
