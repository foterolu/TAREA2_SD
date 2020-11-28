package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"time"

	protos "../protos"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := protos.NewChunksUploadClient(conn)

	fileToBeChunked := "./Orgullo_y_prejuicio-Jane_Austen.pdf" // change here!

	file, err := os.Open(fileToBeChunked)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 0.25 * (1 << 20) // 0.25 MB, change this to your requirement

	// calculate total number of parts the file will be chunked into

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)
	stream, err := client.UploadChunk(ctx)
	if err != nil {

		return
	}
	defer stream.CloseSend()

	for i := uint64(0); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		// write to disk
		fileName := "Orgullo_y_prejuicio-Jane_Austen_" + strconv.FormatUint(i, 10)
		_, err := os.Create(fileName)

		file.Read(partBuffer)
		stream.Send(&protos.Chunk{
			Content: partBuffer,
			Name:    fileName,
		})

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// write/save buffer to disk/send buffer

		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)

		fmt.Println("Split to : ", fileName)
	}

	// just for fun, let's recombine back the chunked files in a new file

	file.Close()

}
