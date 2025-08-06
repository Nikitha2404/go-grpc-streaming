package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Nikitha2404/server-side-streaming/protogen/golang/streamingData"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var SERVER = "0.0.0.0:50051"

func main() {
	conn, err := grpc.NewClient(SERVER, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to streaming service %v", err)
	}
	defer conn.Close()

	client := streamingData.NewStreamingServiceClient(conn)

	req := &streamingData.DataRequest{
		Id: "123",
	}
	stream, err := client.GetDataStreaming(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get response from server %v", err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return
		} else if err != nil {
			log.Fatalf("could not receive server stream %v", err)
		} else {
			valStr := fmt.Sprintf("Response\n Part: %d \n Val: %s", resp.Part, resp.Buffer)
			log.Println(valStr)
		}
	}
}
