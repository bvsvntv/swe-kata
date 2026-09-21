function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

export async function GET() {
  const encoder = new TextEncoder()

  const stream = new ReadableStream({
    async start(controller) {
      try {
        for (let i = 0; i < 10; i++) {
          controller.enqueue(
            encoder.encode(JSON.stringify({ count: i }) + "\n")
          )

          await sleep(500)
        }

        controller.close()
      } catch (error) {
        controller.error(error)
      }
    },
  })

  return new Response(stream, {
    headers: {
      "Content-Type": "application/x-ndjson; charset=utf-8",
      "Cache-Control": "no-cache, no-transform",
    },
  })
}
