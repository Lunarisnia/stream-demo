client:
	go run ./cmd/client

server:
	go run ./cmd/server

generate:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    chatservice/chatservice.proto
