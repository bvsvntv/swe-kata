export type Product = {
  id: number
  internalId: number
  name: string
  description: string
  brand: string
  category: string
  price: number
  currency: string
  stock: number
  ean: string
  color: string
  size: string
  availability: string
}

export type SearchResult = {
  source: string
  products: Product[]
  latency: number | null
  count: number | null
  error: string | null
}
