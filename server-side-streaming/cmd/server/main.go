package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/Nikitha2404/server-side-streaming/protogen/golang/streamingData"
	"google.golang.org/grpc"
)

type StreamingService struct {
	streamingData.UnimplementedStreamingServiceServer
}

func (ss *StreamingService) GetDataStreaming(req *streamingData.DataRequest, srv streamingData.StreamingService_GetDataStreamingServer) error {
	log.Println("Fetch data streaming")

	for i := 0; i <= 5; i++ {
		value := randStringBytes(500)

		resp := streamingData.DataResponse{
			Buffer: value,
			Part:   int32(i),
		}

		if err := srv.Send(&resp); err != nil {
			log.Fatalf("failed to send part:%d error=%v", i, err)
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

func randStringBytes(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyz"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to start server %v", err)
	}
	s := grpc.NewServer()
	streamingData.RegisterStreamingServiceServer(s, &StreamingService{})

	log.Printf("server started at port %v", listener.Addr())

	if err = s.Serve(listener); err != nil {
		log.Fatalf("connection closed abruptly %v", err)
	}
}
