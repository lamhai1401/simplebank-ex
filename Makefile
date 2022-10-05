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

.PHONY: test