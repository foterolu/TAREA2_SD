package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	protos "../protos"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type DataNodeServer struct {
	protos.UnimplementedChunksUploadServer
	data [][]byte
	name []string
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func repartir(dirs []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ip := &protos.Prop{}
	for {

		if len(dirs) == 0 {
			fmt.Printf("No hay nodos funcionando, Todo se guara aca ctm lol xd\n")
			return
		}
		for i := int(0); i < len(dirs); i++ {
			ip.Node = dirs[i]
			aceptacion, err := Propuesta(ctx, ip)
			if !aceptacion.Flag {
				fmt.Printf("Propuesta Rechazada, generando nueva propuesta\n")
				remove(dirs, i)
				break
			}
		}
		/*
			for i := int(0); i < len(DataNodeServer.data); i++ {
				size := len(dirs)

				conn, err := grpc.Dial(dirs[i%size], grpc.WithInsecure())
				if err != nil {
					panic(err)
				}
				defer conn.Close()

				client := protos.NewChunksUploadClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				stream, err := client.UploadChunk(ctx)
				if err != nil {

					return
				}
				defer stream.CloseSend()

				stream.Send(&protos.Chunk{
					Content: DataNodeServer.data[i],
					Name:    DataNodeServer.name[i],
				})

			}

			status, err := stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("search error: %v", err)
				return
			}

			if status.Code != protos.UploadStatusCode_Ok {
				log.Fatalf("search error: %v", err)
				return
			}
		*/

	}
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

	/*
	   ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	   	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure()) //deberia conectarse a cualquiera de los 3 nodeos
	   	if err != nil {
	   		panic(err)
	   	}
	   	defer conn.Close()

	   	if s.data*/
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	defer grpcServer.Stop()
}

func (c *DataNodeServer) Propuesta(ctx context.Context, direccion *protos.Prop) (*protos.Accept, error) { //recibe una ip a confirmar

	aceptacion := &protos.Accept{
		Flag: true}
	conn, err := grpc.Dial(direccion.Node, grpc.WithInsecure()) //deberia conectarse a cualquiera de los 3 nodeos

	if err != nil {
		aceptacion.Flag = false
		panic(err)
		return aceptacion, err
	}
	defer conn.Close()

	return aceptacion, nil

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
			break

		}
		if err != nil {
			log.Fatalf("search error: %v", err)

		}
		//fmt.Printf("Status: %v\n", res.Name)
		s.data = append(s.data, res.Content)
		s.name = append(s.name, res.Name)
		//fmt.Printf("data length: %v\n", len(s.data[len(s.data)-1]))

		ioutil.WriteFile(res.Name, res.Content, os.ModeAppend)
	}
	/*
		for i := int(0); i < len(s.name); i++ {
			fmt.Printf("Nombre : %v\n", s.name[i])
		}

		conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		ctx := stream.Context()
		client := protos.NewChunksUploadClient(conn)

		stream2, err := client.SendChunk(ctx)
		if err != nil {

			return
		}
		defer stream2.CloseSend()

		stream2.Send(&protos.Chunk{
			Content: s.data[0],
			Name:    s.name[0],
		})

		status, err := stream2.CloseAndRecv()
		if err != nil {
			log.Fatalf("chupalo: %v", err)
			return
		}

		if status.Code != protos.UploadStatusCode_Ok {
			log.Fatalf("search error: %v", err)
			return
		}

		return nil*/

	return

}

func (s *DataNodeServer) SendChunk(stream protos.ChunksUpload_SendChunkServer) (err error) {

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

	ioutil.WriteFile(res.Name, res.Content, os.ModeAppend)

	return err
}
