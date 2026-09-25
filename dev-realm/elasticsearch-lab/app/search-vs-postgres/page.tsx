"use client"

import { HugeiconsIcon } from "@hugeicons/react"
import { Search01Icon } from "@hugeicons/core-free-icons"
import { useState } from "react"
import {
  InputGroup,
  InputGroupAddon,
  InputGroupInput,
} from "@/components/ui/input-group"
import { Button } from "@/components/ui/button"
import { SearchResult } from "../types"
import ResultSection from "../ui/result-section"

export default function Page() {
  const [searchResults, setSearchResults] = useState<SearchResult[]>([])
  const [query, setQuery] = useState<string>("")

  const pgResults = searchResults.find(
    (product: SearchResult) => product.source === "PostgreSQL (ILIKE)"
  )

  const esResults = searchResults.find(
    (product: SearchResult) => product.source === "Elasticsearch"
  )

  async function handleSearch() {
    if (!query.trim()) return

    setSearchResults([])

    const response = await fetch(`/api/products?q=${encodeURIComponent(query)}`)

    const reader = response?.body!.getReader()
    const decoder = new TextDecoder()

    let buffer = ""
    while (true) {
      const { value, done } = await reader.read()
      if (done) break

      buffer += decoder.decode(value, { stream: true })
      const lines = buffer.split("\n")
      buffer = lines.pop() ?? ""

      for (const line of lines) {
        if (!line.trim()) continue

        const result = JSON.parse(line) as SearchResult
        setSearchResults((prev) => [...prev, result])
      }

      // Flush TextDecoder
      buffer += decoder.decode()

      if (buffer.trim()) {
        const result = JSON.parse(buffer)
        console.log(result)
      }
    }
  }

  return (
    <main className="min-h-screen bg-muted/30 px-4 py-8 sm:px-8">
      <div className="mx-auto max-w-6xl">
        <div className="mb-6 space-y-1">
          <h1 className="font-heading text-lg font-semibold tracking-tight">
            Product search comparison
          </h1>
          <p className="text-sm text-muted-foreground">
            Compare PostgreSQL substring search with Elasticsearch.
          </p>
        </div>

        <InputGroup className="max-w-2xl">
          <InputGroupInput
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                handleSearch()
              }
            }}
            onChange={(e) => setQuery(e.target.value)}
            placeholder="Search..."
          />
          <InputGroupAddon>
            <Button
              aria-label="Search products"
              onClick={handleSearch}
              size="icon-sm"
              variant="ghost"
            >
              <HugeiconsIcon icon={Search01Icon} strokeWidth={1.5} />
            </Button>
          </InputGroupAddon>
        </InputGroup>

        <div className="mt-6 grid gap-4 lg:grid-cols-2">
          <ResultSection title="PostgreSQL Search" searchResult={pgResults} />

          <ResultSection title="Elasticsearch" searchResult={esResults} />
        </div>
      </div>
    </main>
  )
}
