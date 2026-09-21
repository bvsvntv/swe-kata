"use client"

export default function Page() {
  async function start() {
    const response = await fetch("/api/stream")

    const reader = response.body!.getReader()
    const decoder = new TextDecoder()

    while (true) {
      const { value, done } = await reader.read()
      if (done) break

      console.log("stream result: ", decoder.decode(value))
    }
  }

  return (
    <button onClick={start} className="m-4 rounded-md bg-gray-500 p-2">
      Start streaming
    </button>
  )
}
