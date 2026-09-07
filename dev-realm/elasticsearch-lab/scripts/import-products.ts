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
  let records
  try {
    const productsFilePath = path.join(process.cwd(), "products-2000000.csv")
    if (!fs.existsSync(productsFilePath)) {
      console.log(
        `ERROR: CSV file not found at ${productsFilePath}. Make sure the path is correct relative to the project root`
      )
      process.exit(1)
    }

    const fileContent = fs.readFileSync(productsFilePath, { encoding: "utf-8" })
    records = parse(fileContent, {
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
    console.log("Importing records to postgresql...")

    // Ensure that table exists
    await prisma.$executeRaw`
  CREATE TABLE IF NOT EXISTS "products" (
    "id" SERIAL NOT NULL,
    "internal_id" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "brand" TEXT NOT NULL,
    "category" TEXT NOT NULL,
    "price" DECIMAL(10,2) NOT NULL,
    "currency" TEXT NOT NULL,
    "stock" INTEGER NOT NULL,
    "ean" TEXT NOT NULL,
    "color" TEXT NOT NULL,
    "size" TEXT NOT NULL,
    "availability" TEXT NOT NULL,

    CONSTRAINT "products_pkey" PRIMARY KEY ("id")
  );

  CREATE UNIQUE INDEX IF NOT EXISTS "products_ean_key"
    ON "products"("ean");

  CREATE INDEX IF NOT EXISTS "products_name_idx"
    ON "products"("name");

  CREATE INDEX IF NOT EXISTS "products_brand_idx"
    ON "products"("brand");

  CREATE INDEX IF NOT EXISTS "products_category_idx"
    ON "products"("category");

  CREATE INDEX IF NOT EXISTS "products_availability_idx"
    ON "products"("availability");

  CREATE INDEX IF NOT EXISTS "products_price_idx"
    ON "products"("price");

  CREATE INDEX IF NOT EXISTS "products_stock_idx"
    ON "products"("stock");
`

    await prisma.$executeRaw`TRUNCATE TABLE products RESTART IDENTITY CASCADE;`
    console.log("Truncated existing data from postgres 'products' table.")

    // Validate records
    const validRecords = records
      .filter(
        (record: any) =>
          record["Internal ID"] &&
          record.Name &&
          record.Description &&
          record.Brand &&
          record.Category &&
          record.Price &&
          record.Currency &&
          record.Stock &&
          record.EAN &&
          record.Color &&
          record.Size &&
          record.Availability
      )
      .map((record: any) => ({
        internalId: Number(record["Internal ID"]),
        name: record.Name,
        description: record.Description,
        brand: record.Brand,
        category: record.Category,
        price: Number(record.Price),
        currency: record.Currency,
        stock: Number(record.Stock),
        ean: record.EAN,
        color: record.Color,
        size: record.Size,
        availability: record.Availability,
      }))

    if (validRecords.length > 0) {
      const BATCH_SIZE = 1000
      let totalInserted = 0

      for (let i = 0; i <= validRecords.length; i += BATCH_SIZE) {
        const batch = validRecords.slice(i, i + BATCH_SIZE)
        const result = await prisma.product.createMany({
          data: batch,
        })

        totalInserted += result.count
        console.log(
          `Inserted ${totalInserted} / ${validRecords.length} products`
        )
      }

      console.log(`Saved ${totalInserted} products to database.`)
    }
  } catch (e) {
    console.log("ERROR: Failed to import records into table. Error: ", e)
    process.exit(1)
  }
}

await main()
