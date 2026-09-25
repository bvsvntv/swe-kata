"use client"

import { buttonVariants } from "@/components/ui/button"

export default function Page() {
  return (
    <main className="min-h-screen bg-muted/30 px-4 py-8 sm:px-8">
      <div className="mx-auto max-w-6xl">
        <div className="flex flex-col items-start gap-8 sm:flex-row">
          <a
            href="/search-vs-postgres"
            className={buttonVariants({ variant: "secondary", size: "sm" })}
          >
            postgresql-substring-search-vs-elasticsearch
          </a>
        </div>
      </div>
    </main>
  )
}
