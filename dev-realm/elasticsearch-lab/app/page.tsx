"use client"

import { useState } from "react"

export default function Page() {
  const [text, setText] = useState("")

  async function start() {
    const response = await fetch("/api/stream")

    const reader = response.body!.getReader()
    const decoder = new TextDecoder()

    while (true) {
      const { value, done } = await reader.read()
      if (done) break

      setText(decoder.decode(value))
    }
  }

  return (
    <div className="m-4">
      <button onClick={start} className="rounded-md bg-gray-500 p-2">
        Start streaming
      </button>

      {text ? (
        <p className="mt-4 text-gray-500">{text}</p>
      ) : (
        <p>Not streaming.</p>
      )}
    </div>
  )
}
