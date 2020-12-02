package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	protos "../protos"
	"google.golang.org/grpc"
)

const (
	node0    = "localhost:5050"
	node3    = "localhost:9090"
	node1    = "localhost:50051"
	node2    = "localhost:8080"
	namenode = "localhost:4040"
)

var (
	cliente = flag.String("cliente", "", "tipo cliente")
	libro   = flag.String("libro", "", "nombre del libro")
	mu      sync.RWMutex
)

//DataNodeServer Estructura del server
type DataNodeServer struct {
	protos.UnimplementedChunksUploadServer
	chunk []*protos.Chunk
	data  [][]byte
	name  []string
	dir   []string
}

func main() {
	flag.Parse()
	if *cliente == "uploader" {
		conn, err := grpc.Dial(node1, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second)) //deberia conectarse a cualquiera de los 3 nodeos

		if err != nil {
			fmt.Printf("ERRROOOOOOR\n")
			panic(err)
		}
		defer conn.Close()

		client := protos.NewChunksUploadClient(conn)

		fmt.Printf("nombre del libro %v\n", *libro)

		fileToBeChunked := "./upload/" + *libro + ".pdf" // change here!

		file, err := os.Open(fileToBeChunked)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer file.Close() //pa que termine despues de ejecutarse todo

		fileInfo, _ := file.Stat()

		var fileSize int64 = fileInfo.Size()

		const fileChunk = 0.25 * (1 << 20) // 0.25 MB, change this to your requirement

		// calculate total number of parts the file will be chunked into

		totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

		fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
		//ctx := context.Background()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		stream, err := client.UploadChunk(ctx)
		if err != nil {

			return
		}
		defer stream.CloseSend()
		//fmt.Printf("Total part: %v\n", totalPartsNum)

		for i := uint64(0); i < totalPartsNum; i++ {

			partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
			partBuffer := make([]byte, partSize)

			// write to disk
			fileName := *libro + "_" + strconv.FormatUint(i, 10)
			//_, err := os.Create(fileName)

			file.Read(partBuffer)
			fmt.Printf("cantidad de bytes: %v\n", len(partBuffer))
			stream.Send(&protos.Chunk{
				Content: partBuffer,
				Name:    fileName,
				Libro:   *libro,
				Partes:  int32(totalPartsNum),
			})

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// write/save buffer to disk/send buffer

			//ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

			fmt.Println("Split to : ", fileName)
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

		// just for fun, let's recombine back the chunked files in a new file

		//file.Close()
	}

	if *cliente == "downloader" {
		conn, err := grpc.Dial(namenode, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
		if err != nil {
			fmt.Printf("ERRROOOOOOR\n")
			panic(err)
		}
		defer conn.Close()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client := protos.NewChunksUploadClient(conn)
		p := &protos.Prop{
			Node: *libro,
		}
		mu.Lock()

		ad, _ := client.RequestAdress(ctx, p)

		// Inicializacion del server
		/*
			listener, err := net.Listen("tcp", node0)
			if err != nil {
				panic(err)
			}
			var opts []grpc.ServerOption
			grpcServer := grpc.NewServer(opts...)

			s := &DataNodeServer{}
			protos.RegisterChunksUploadServer(grpcServer, s)
			fmt.Printf("escuchando\n")
			grpcServer.Serve(listener)



			defer grpcServer.Stop()
		*/
		newFileName := "downloads/" + *libro + ".pdf"
		_, err = os.Create(newFileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		file, err := os.OpenFile(newFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//var writePosition int64 = 0

		//Aca crear metodos para recibir chunks

		for i := 0; i < len(ad.Adress); i++ {
			data := strings.Split(ad.Adress[i], " ")

			conn2, err := grpc.Dial(data[1], grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
			if err != nil {
				fmt.Printf("ERRROOOOOOR\n")
				panic(err)
			}
			defer conn.Close()
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			client2 := protos.NewChunksUploadClient(conn2)
			pr := &protos.Prop{
				Node: data[0]}

			chunk, err := client2.DownloadChunk(ctx, pr)
			fmt.Printf("%v\n", chunk)
			chunkBufferBytes := chunk.Content
			//writePosition = writePosition + chunk.Partes

			n, err := file.Write(chunkBufferBytes)
			file.Sync() //flush to disk

			chunkBufferBytes = nil
			fmt.Println("Written ", n, " bytes")

		}
		file.Close()
		mu.Unlock()

	}

}
