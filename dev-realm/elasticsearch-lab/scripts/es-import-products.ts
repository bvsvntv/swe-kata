import dotenv from "dotenv"
import { Client } from "@elastic/elasticsearch"
import { HttpConnection } from "@elastic/transport"
import { readCSV } from "@/utils/read-csv.util"

dotenv.config({ path: "./.env" })

async function main() {
  const ELASTICSEARCH_INDEX_NAME = "products"

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
    const esIndexExists = await esClient.indices.exists({
      index: ELASTICSEARCH_INDEX_NAME,
    })
    if (esIndexExists) {
      console.log(`Elasticsearch index '${ELASTICSEARCH_INDEX_NAME}' exists.`)
      await esClient.indices.delete({ index: ELASTICSEARCH_INDEX_NAME })
      console.log(`Elasticsearch index '${ELASTICSEARCH_INDEX_NAME}' deleted.`)
    }

    await esClient.indices.create({
      index: ELASTICSEARCH_INDEX_NAME,
      settings: {
        number_of_shards: 1,
        number_of_replicas: 1,
      },
      mappings: {
        properties: {
          id: { type: "integer" },
          internalId: { type: "integer" },

          name: {
            type: "text",
            fields: {
              keyword: { type: "keyword", ignore_above: 256 },
            },
          },
          description: { type: "text" },
          brand: { type: "keyword" },
          category: { type: "keyword" },
          color: { type: "keyword" },
          size: { type: "keyword" },
          availability: { type: "keyword" },
          price: { type: "scaled_float", scaling_factor: 100 },
          currency: { type: "keyword" },
          stock: { type: "integer" },
          ean: { type: "keyword" },
        },
      },
    })
    console.log(
      `New elasticsearch index '${ELASTICSEARCH_INDEX_NAME}' created.`
    )

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
