package main

import (
	"fmt"
	"net"

	protos "../protos"
	"google.golang.org/grpc"
)

const (
	dir0 = "localhost:9090"
	dir1 = "localhost:50051"
	dir2 = "localhost:8080"
	dir3 = "localhost:4040"
)

type NameNodeServer struct {
	protos.UnimplementedChunksUploadServer
}

func main() {

	listener, err := net.Listen("tcp", dir0)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	protos.RegisterChunksUploadServer(grpcServer, &NameNodeServer{})
	fmt.Printf("escuchando\n")
	grpcServer.Serve(listener)

	defer grpcServer.Stop()
}
