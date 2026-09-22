"use client"

import { HugeiconsIcon } from "@hugeicons/react"
import { Search01Icon } from "@hugeicons/core-free-icons"
import { useState } from "react"
import { SearchResult } from "./types"
import ResultSection from "./ui/result-section"
import {
  InputGroup,
  InputGroupAddon,
  InputGroupInput,
} from "@/components/ui/input-group"

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
    <div className="m-8 mx-auto max-w-5xl">
      <InputGroup className="max-w-5xl">
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
          <HugeiconsIcon
            icon={Search01Icon}
            size={24}
            color="currentColor"
            onClick={handleSearch}
            strokeWidth={1.5}
          />
        </InputGroupAddon>
      </InputGroup>

      <div className="flex gap-4">
        <div className="min-w-0 flex-1">
          <ResultSection
            title="PostgreSQL Search"
            searchResult={pgResults as SearchResult}
          />
        </div>

        <div className="min-w-0 flex-1">
          <ResultSection
            title="Elasticsearch"
            searchResult={esResults as SearchResult}
          />
        </div>
      </div>
    </div>
  )
}
