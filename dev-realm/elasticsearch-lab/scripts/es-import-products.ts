import dotenv from "dotenv"
import { Client } from "@elastic/elasticsearch"
import { HttpConnection } from "@elastic/transport"

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
}

await main()
