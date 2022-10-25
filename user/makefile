proto:
	protoc ./proto/*.proto \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt paths=source_relative \
		--proto_path=.

.PHONY: proto