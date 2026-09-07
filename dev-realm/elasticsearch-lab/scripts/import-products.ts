import fs from "node:fs"
import dotenv from "dotenv"
import { PrismaPg } from "@prisma/adapter-pg"
import { PrismaClient } from "@/generated/prisma/client"
import path from "node:path"
import { parse } from "csv-parse/sync"

dotenv.config({ path: "./.env" })

async function main() {
  const dbURL = process.env.DATABASE_URL
  if (!dbURL) {
    console.log("ERROR: DATABASE_URL missing from '.env' file.")
    process.exit(1)
  }

  // prisma client
  const adapter = new PrismaPg({ connectionString: dbURL })
  const prisma = new PrismaClient({ adapter })

  try {
    await prisma.$connect()
    console.log("Prisma client initialized.")
  } catch (e) {
    console.log("ERROR: Failed to initialize prisma client. Error: ", e)
    process.exit(1)
  }

  // read csv
  try {
    const productsFilePath = path.join(process.cwd(), "products-2000000.csv")
    if (!fs.existsSync(productsFilePath)) {
      console.log(
        `ERROR: CSV file not found at ${productsFilePath}. Make sure the path is correct relative to the project root`
      )
      process.exit(1)
    }

    const fileContent = fs.readFileSync(productsFilePath, { encoding: "utf-8" })
    const records = parse(fileContent, {
      columns: true,
      skip_empty_lines: true,
    })

    console.log(`Read ${records.length} records from CSV.`)
  } catch (e) {
    console.log("ERROR: Failed to read or parse csv file. Error: ", e)
    process.exit(1)
  }

  // Import records into table
  try {
  } catch (e) {
    console.log("ERROR: Failed to import records into table. Error: ", e)
    process.exit(1)
  }
}

await main()
