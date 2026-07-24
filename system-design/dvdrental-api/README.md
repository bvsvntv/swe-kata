# dvdrental-api

## Notes

- Extract db schema only:

  ```sh
  PGPASSWORD='postgres' \
  pg_dump \
      -U postgres \
      -h localhost \
      -p 5432 \
      --schema-only \
      --no-owner \
      --no-privileges \
      dvdrental > schema.sql
  ```

## To-Do
