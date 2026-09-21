function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}

export async function GET() {
  const encoder = new TextEncoder()
  const stream = new ReadableStream({
    async start(controller) {
      try {
        for (let i = 0; i <= 10; i++) {
          controller.enqueue(encoder.encode(`Processing step: ${i}\n`))

          await sleep(1000)
        }

        controller.close()
      } catch (error: unknown) {
        controller.error(error)
      }
    },
  })

  return new Response(stream, {
    headers: {
      "Content-Type": "text/plain; charset=utf-8",
    },
  })
}
