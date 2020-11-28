package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

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

func (s *DataNodeServer) UploadChunk(stream protos.ChunksUpload_UploadChunkServer) (err error) {
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			err = stream.SendAndClose(&protos.UploadStatus{
				Message: "Upload received with success",
				Code:    protos.UploadStatusCode_Ok,
			})
			if err != nil {
				log.Fatalf("search error: %v", err)
				return err
			}
			return nil

		}
		if err != nil {
			log.Fatalf("search error: %v", err)

		}
		fmt.Printf("Status: %v\n", res.Name)

		ioutil.WriteFile(res.Name, res.Content, os.ModeAppend)

	}
	//return err

}
