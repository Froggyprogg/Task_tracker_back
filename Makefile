PHONY: generate-structs
generate-structs:
	IF exist "pkg/user_v1" ( rmdir /s "pkg/user_v1" )


	mkdir "pkg/user_v1"
	protoc --go_out=pkg/user_v1 --go_opt=paths=import \
			--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=import \
			api/user_v1/service.proto

PHONY: generate-board
generate-board:
	IF exist "pkg/board_v1" ( rmdir /s "pkg/board_v1" )


	mkdir "pkg/board_v1"
	protoc --go_out=pkg/board_v1 --go_opt=paths=import \
			--go-grpc_out=pkg/board_v1 --go-grpc_opt=paths=import \
			api/board_v1/service.proto
