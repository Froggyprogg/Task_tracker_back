PHONY: generate-structs
generate-structs:
	mkdir "pkg/user_v1"
	protoc --go_out=pkg/user_v1 --go_opt=paths=import \
			--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=import \
			api/user_v1/service.proto
