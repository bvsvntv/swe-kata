import fs from "node:fs"
import path from "node:path"
import { parse } from "csv-parse/sync"

export function readCSV(fileName: string) {
  try {
    const file = path.join(process.cwd(), fileName)
    if (!fs.existsSync(file)) {
      console.log(
        `ERROR: CSV file not found at ${file}. Make sure the path is correct relative to the project root`
      )
      process.exit(1)
    }

    const fileContent = fs.readFileSync(file, { encoding: "utf-8" })
    const records = parse(fileContent, {
      columns: true,
      skip_empty_lines: true,
    })

    console.log(`Read ${records.length} records from CSV.`)
    return records
  } catch (e) {
    console.log("ERROR: Failed to read or parse csv file. Error: ", e)
    process.exit(1)
  }
}
