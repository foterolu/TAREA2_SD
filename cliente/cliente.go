package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"io/ioutil"

	protos "../protos"
	"google.golang.org/grpc"
)

const (
	node3    = "10.10.28.47:8080"
	node1    = "10.10.28.45:8080"
	node2    = "10.10.28.46:8080"
	namenode = "10.10.28.48:8080"
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
		directorio, err := ioutil.ReadDir("./upload")
   		if err != nil {
			panic(err)
    			}
		fmt.Printf("%v\n", directorio[0].Name()) 
		cantidadLibros := len(directorio)
		numeroRandom := rand.Intn(cantidadLibros)
		
		conn, err := grpc.Dial(node1, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(25*time.Second)) //deberia conectarse a cualquiera de los 3 nodeos

		if err != nil {
			fmt.Printf("ERRROOOOOOR\n")
			panic(err)
		}
		defer conn.Close()

		client := protos.NewChunksUploadClient(conn)

		fmt.Printf("nombre del libro %v\n", directorio[numeroRandom].Name())

		fileToBeChunked := "./upload/" + directorio[numeroRandom].Name() // change here!

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

		ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
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
