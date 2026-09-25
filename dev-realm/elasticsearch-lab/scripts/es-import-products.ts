import dotenv from "dotenv"
import { Client } from "@elastic/elasticsearch"
import { HttpConnection } from "@elastic/transport"
import { readCSV } from "@/utils/read-csv.util"

dotenv.config({ path: "./.env" })

type CsvRecord = Record<string, string>

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

  const records = readCSV("products-100000.csv") as CsvRecord[]

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
        number_of_replicas: 0,
      },
      mappings: {
        properties: {
          id: { type: "integer" },
          internalId: { type: "integer" },
          name: {
            type: "text",
            fields: {
              raw: { type: "keyword" },
            },
          },
          description: {
            type: "text",
            fields: {
              raw: { type: "keyword" },
            },
          },
          brand: {
            type: "text",
            fields: {
              raw: { type: "keyword" },
            },
          },
          category: {
            type: "text",
            fields: {
              raw: { type: "keyword" },
            },
          },
          color: {
            type: "text",
            fields: {
              raw: { type: "keyword" },
            },
          },
          size: {
            type: "text",
            fields: {
              raw: { type: "keyword" },
            },
          },
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
        (record) =>
          record.Index &&
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
      .map((record) => ({
        id: Number(record.Index),
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
      console.log(
        `${validRecords.length} records are valid. Indexing them in elasticsearch.`
      )
      const BATCH_SIZE = 10000
      let totalInserted = 0

      for (let i = 0; i < validRecords.length; i += BATCH_SIZE) {
        const batch = validRecords.slice(i, i + BATCH_SIZE)
        const bulkBody = batch.flatMap((record) => [
          {
            index: {
              _index: ELASTICSEARCH_INDEX_NAME,
              _id: record.id,
            },
          },
          record,
        ])

        const bulkResponse = await esClient.bulk({
          refresh: false,
          body: bulkBody,
        })

        if (bulkResponse.errors) {
          console.error("Some records failed to index while bulk indexing.")

          for (const item of bulkResponse.items) {
            if (item.index && item.index.error) {
              console.error(
                `Error indexing document ID ${item.index._id}: ${JSON.stringify(item.index.error)}`
              )
            }
          }

          throw new Error(`Bulk indexing failed for batch starting at ${i}.`)
        } else {
          console.log(`${batch.length} records imported successfully.`)
        }

        totalInserted += bulkResponse.items.filter(
          (item) =>
            item.index?.result === "created" || item.index?.result === "updated"
        ).length
        console.log(
          `Indexed ${totalInserted} / ${validRecords.length} products`
        )
      }

      // Only refresh index after full insert
      await esClient.indices.refresh({
        index: ELASTICSEARCH_INDEX_NAME,
      })

      const indexedCount = await esClient.count({
        index: ELASTICSEARCH_INDEX_NAME,
      })
      if (indexedCount.count !== validRecords.length) {
        throw new Error(
          `Indexed ${indexedCount.count} documents, expected ${validRecords.length}.`
        )
      }
      console.log(`Verified ${indexedCount.count} indexed products.`)
    }
  } catch (e) {
    console.log(
      "ERROR: Failed to import records into elasticsearch index. Error: ",
      e
    )
    process.exit(1)
  } finally {
    await esClient.close()
    console.log("Elasticsearch client closed.")
  }
}

await main()
