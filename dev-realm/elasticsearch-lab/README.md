# elasticsearch-lab

## Notes

### Resources

- [ 15. Full text search using Elasticsearch for blazingly fast search ](https://www.youtube.com/watch?v=7_sovzAhRSM&t=1747s)

### Getting Started

- Download the CSV records from [here](https://www.datablist.com/learn/csv/download-sample-csv-files)
- Seed them into PostgreSQL and Elasticsearch
  ```sh
  bun run pg:import-products
  bun run es:import-products
  ```
- Visit [http://localhost:3000](http://localhost:3000/) and perform search

## To-Do

- [x] Compare substring search in PostgreSQL and Elasticsearch
- [ ] Autocomplete suggestion from Elasticsearch
