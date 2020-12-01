package main

import (
	"context"
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
	diccionario map[string][]string
}

func main() {

	listener, err := net.Listen("tcp", dir3)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	s := &NameNodeServer{}

	s.diccionario = make(map[string][]string)
	protos.RegisterChunksUploadServer(grpcServer, s)
	fmt.Printf("escuchando\n")
	grpcServer.Serve(listener)
	defer grpcServer.Stop()

}

func (s *NameNodeServer) SendLog(ctx context.Context, report *protos.Log) (*protos.Accept, error) {
	accept := &protos.Accept{}
	fmt.Printf("QUE RECHUCHA PIJA PICO\n")
	s.diccionario[report.NombreLibro+" "+report.CantidadPartes] = append(s.diccionario[report.NombreLibro+" "+report.CantidadPartes], report.Parte+" "+report.Ubicaciones)
	fmt.Printf("Diccionario %v \n", s.diccionario)
	return accept, nil

}
