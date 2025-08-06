# Server side streaming

##### Run following command to create protobug file
    cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative \
	--go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative \
	./**/*.proto

##### Run server
    cd cmd/server
    go run main.go

##### Run client
    cd cmd/client
    go run main.go