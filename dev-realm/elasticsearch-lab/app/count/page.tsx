"use client"
import { useState } from "react"

export default function Page() {
  const [count, setCount] = useState(0)

  async function startCount() {
    const response = await fetch("/api/count")

    const reader = response?.body!.getReader()
    const decoder = new TextDecoder()

    while (true) {
      const { value, done } = await reader.read()
      if (done) break

      setCount(Number(decoder.decode(value)))
    }
  }

  return (
    <div className="m-4">
      <button onClick={startCount}>start counting</button>

      <p>{count}</p>
    </div>
  )
}
