package main

import (
	"context"
	"fmt"
	"net"

	protos "../protos"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type DataNodeServer struct {
	protos.UnimplementedChunksUploadServer
}

func main() {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	protos.RegisterChunksUploadServer(grpcServer, &DataNodeServer{})
	fmt.Printf("escuchando\n")
	grpcServer.Serve(listener)

	defer grpcServer.Stop()
}

func (s *DataNodeServer) UploadChunk(ctx context.Context, Chunk *protos.Chunk) (*protos.UploadStatus, error) {

	fmt.Printf("Status: %v", Chunk.Content)
	return &protos.UploadStatus{}, nil

}
