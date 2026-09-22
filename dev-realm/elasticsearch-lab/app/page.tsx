"use client"

export default function Page() {
  async function handleSearch() {
    const response = await fetch("/api/products")

    const reader = response?.body!.getReader()
    const decoder = new TextDecoder()

    let buffer = ""
    while (true) {
      const { value, done } = await reader.read()
      if (done) break

      buffer += decoder.decode(value, { stream: true })
      const lines = buffer.split("\n")
      buffer = lines.pop() ?? ""

      lines.forEach((line) => {
        const result = JSON.parse(line)
        console.log(result)
      })
    }
  }

  return (
    <div className="m-4">
      <button onClick={handleSearch}>search products</button>
    </div>
  )
}
