"use client"
import { useState } from "react"

export default function Page() {
  const [count, setCount] = useState(0)

  async function startCount() {
    const response = await fetch("/api/count")

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
        setCount(result.count)
      })
    }
  }

  return (
    <div className="m-4">
      <button onClick={startCount}>start counting</button>

      <p>{count}</p>
    </div>
  )
}
