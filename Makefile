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
	migrate -path ./db/migration -database "postgresql://root:secret@localhosT:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhosT:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: test