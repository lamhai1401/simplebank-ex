# simplebank-ex
How to build a simple bank with golang and gRPC, microservices

## Flowing

### Init schema

migrate create -ext sql -dir /db/migration -seq init_schema

### Run pg db

make pg-start

#### init account

createdb --username=root --owner=root simple_bank
dropdb simple_bank
psql simple_bank

### Migrate PG db

migrate -path ./db/migration -database "postgresql://root:secret@localhosT:5432/simple_bank?sslmode=disable" -verbose up
