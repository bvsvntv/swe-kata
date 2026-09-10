import dotenv from "dotenv"
import { Client } from "@elastic/elasticsearch"
import { HttpConnection } from "@elastic/transport"
import { readCSV } from "@/utils/read-csv.util"

dotenv.config({ path: "./.env" })

async function main() {
  const esNodeURL = process.env.ELASTICSEARCH_NODE_URL
  if (!esNodeURL) {
    console.log("ERROR: ELASTICSEARCH_NODE_URL missing from '.env' file.")
    process.exit(1)
  }

  const esClient = new Client({ node: esNodeURL, Connection: HttpConnection })

  try {
    await esClient.ping()
    console.log("Elasticsearch client initialized.")
  } catch (e) {
    console.log("ERROR: Failed to initialize elasticsearch client. Error: ", e)
    process.exit(1)
  }

  const records = readCSV("products-2000000.csv")

  // Import records into elasticsearch index
  try {
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
      console.log(`${validRecords.length} records are valid.`)
    }
  } catch (e) {
    console.log(
      "ERROR: Failed to import records into elasticsearch index. Error: ",
      e
    )
    process.exit(1)
  }
}

await main()
