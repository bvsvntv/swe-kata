import "dotenv/config"
import { PrismaPg } from "@prisma/adapter-pg"
import { PrismaClient } from "../generated/prisma/client"

const connectionString = `${process.env.DATABASE_URL}`

const adapter = new PrismaPg({ connectionString })
const prisma = new PrismaClient({ adapter })

// handle graceful shutdown
process.on("beforeExit", async () => {
  console.log("> Closing database connection.")
  await prisma.$disconnect()
})

process.on("SIGINT", async () => {
  console.log("> Closing database connection.")
  await prisma.$disconnect()
})

process.on("SIGTERM", async () => {
  console.log("> Closing database connection.")
  await prisma.$disconnect()
})

export { prisma }
