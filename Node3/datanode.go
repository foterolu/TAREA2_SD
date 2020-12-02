package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	protos "../protos"
	"google.golang.org/grpc"
)

const (
	node3    = "localhost:9090"
	node1    = "localhost:50051"
	node2    = "localhost:8080"
	namenode = "localhost:4040"
)

var (
	mu sync.RWMutex
)

type DataNodeServer struct {
	protos.UnimplementedChunksUploadServer
	chunk []*protos.Chunk
	data  [][]byte
	name  []string
	dir   []string
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	listener, err := net.Listen("tcp", node3)
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	di := []string{node3, node2, node1}
	s := &DataNodeServer{}
	s.dir = di
	protos.RegisterChunksUploadServer(grpcServer, s)
	fmt.Printf("escuchando\n")
	grpcServer.Serve(listener)

	fmt.Printf("escuchando\n")

	defer grpcServer.Stop()
}

func (s *DataNodeServer) UploadChunk(stream protos.ChunksUpload_UploadChunkServer) (err error) {

	mu.Lock()
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

		s.chunk = append(s.chunk, res)
		s.data = append(s.data, res.Content)
		s.name = append(s.name, res.Name)
	}
	fmt.Printf("Cantidad de CHunks LOLXD : %v\n", len(s.chunk))
	repartir(s.dir, s)
	mu.Unlock()

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
			//log.Fatalf("Es aacA?: %v", err)
			return err
		}
		return nil

	}
	if err != nil {
		fmt.Printf("UPALE UPSI\n")
		return
		//log.Fatalf("OOOO ES ACA?: %v", err)

	}

	ioutil.WriteFile(res.Name, res.Content, os.ModeAppend)

	return err
}

func (s *DataNodeServer) Propuesta(ctx context.Context, direccion *protos.Prop) (*protos.Accept, error) { //recibe una ip a confirmar

	aceptacion := &protos.Accept{
		Flag: true}
	conn, err := grpc.Dial(direccion.Node, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(1*time.Second)) //deberia conectarse a cualquiera de los 3 nodeos

	if err != nil {
		aceptacion.Flag = false

		return aceptacion, err
	}
	conn.Close()

	return aceptacion, nil

}

func repartir(dirs []string, s *DataNodeServer) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	flag := false
	ip := &protos.Prop{}
	for { //loop infinito

		if len(dirs) == 0 {
			fmt.Printf("No hay nodos funcionando, Todo se guara aca ctm lol xd\n")
			break
		} else if len(dirs) == 1 {
			ip.Node = dirs[0]
			aceptacion, _ := s.Propuesta(ctx, ip)
			fmt.Printf("La propuesta es: %v \n", aceptacion.Flag)
			if !aceptacion.Flag {
				fmt.Printf("Propuesta Rechazada, no Hay nodos disponibles\n")
				dirs = remove(dirs, 0)
				return
			} else {
				break

			}

		}

		for i := int(0); i < len(dirs); i++ {
			if dirs[i] != node3 {

				ip.Node = dirs[i]
				aceptacion, _ := s.Propuesta(ctx, ip)
				fmt.Printf("La propuesta es: %v \n", dirs)
				if !aceptacion.Flag {
					fmt.Printf("Propuesta Rechazada, generando nueva propuesta\n")
					dirs = remove(dirs, i)
					break
				} else if aceptacion.Flag && i == len(dirs)-1 {
					fmt.Printf("Propuesta aceptada: %v \n", dirs)
					flag = true
				}
			}
		}

		if flag {
			break
		}

	}
	fmt.Printf("Cantidad de CHUNKS: %v\n", len(s.data))
	for i := int(0); i < len(s.data); i++ {
		size := len(dirs)
		conn, err := grpc.Dial(dirs[i%size], grpc.WithInsecure())
		if err != nil {

			panic(err)
		}
		defer conn.Close()

		client := protos.NewChunksUploadClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if dirs[i%size] == node3 {
			ioutil.WriteFile(s.name[i], s.data[i], os.ModeAppend)
			reporte := (&protos.Log{
				NombreLibro:    s.chunk[i].Libro,
				CantidadPartes: strconv.FormatInt(int64(s.chunk[i].Partes), 10),
				Ubicaciones:    dirs[i%size],
				Parte:          s.chunk[i].Name,
			})
			conn2, err := grpc.Dial(namenode, grpc.WithInsecure())
			if err != nil {

				panic(err)
			}
			client2 := protos.NewChunksUploadClient(conn2)
			client2.SendLog(ctx, reporte)
			conn2.Close()

			fmt.Printf("Reporte: %v\n", reporte)
		} else {

			stream, err := client.SendChunk(ctx)
			if err != nil {

				return
			}
			defer stream.CloseSend()

			stream.Send(&protos.Chunk{
				Content: s.data[i],
				Name:    s.name[i],
			})

			reporte := (&protos.Log{
				NombreLibro:    s.chunk[i].Libro,
				CantidadPartes: strconv.FormatInt(int64(s.chunk[i].Partes), 10),
				Ubicaciones:    dirs[i%size],
				Parte:          s.chunk[i].Name,
			})
			conn2, err := grpc.Dial(namenode, grpc.WithInsecure())
			if err != nil {

				panic(err)
			}
			client2 := protos.NewChunksUploadClient(conn2)
			client2.SendLog(ctx, reporte)
			conn2.Close()

			client.SendLog(ctx, reporte)

		}

	}
	s.data = nil
	s.name = nil
	s.chunk = nil
	di := []string{node3, node2, node1}
	s.dir = di
	fmt.Printf("data server array: %v", s.data)

}
