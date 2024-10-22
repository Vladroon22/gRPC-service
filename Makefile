.PHONY:

rm-proto:
	sudo rm -rf pkg

grpc:
	mkdir -p pkg/proto
	protoc -I ./internal/proto --go_out=./pkg/proto --go_opt=paths=source_relative --go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative ./internal/proto/add.proto

run_client: 
	go build -o ./cli cmd_client/main.go
	./cli

run_server: 
	go build -o ./srv cmd_server/main.go
	./srv
