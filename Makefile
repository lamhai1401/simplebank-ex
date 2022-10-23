SHELL := /bin/bash

proto:
	protoc ./evenstream/eventstore/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.

swagger:
	swagger generate spec -o ./swagger.json
	swagger serve ./swagger.json

nats:
	docker-compose -f nats-cluster.yml up

nats-stop:
	docker-compose -f nats-cluster.yml down && docker-compose -f nats-cluster.yml stop

init-schema:
	migrate create -ext sql -dir /db/migration -seq init_schema

pg-start:
	docker-compose -f postgres.yml up

pg-stop:
	docker-compose -f postgres.yml down && docker-compose -f postgres.yml stop

migrateup:
	migrate -path ./db/migration -database "postgresql://username:password@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://username:password@localhost:5432/simplebank?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

test-local:
	go test -v -race -coverprofile=c.out -cover ./... ./tests/...
	go tool cover -html=c.out -o coverage.html

.PHONY: test test-local