package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	protos "../protos"
	"google.golang.org/grpc"
)

const (
	dir0 = "localhost:9090"
	dir1 = "localhost:50051"
	dir2 = "localhost:8080"
	dir3 = "localhost:4040"
)

var cont int

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
	cont = 0

	s.diccionario = make(map[string][]string)
	protos.RegisterChunksUploadServer(grpcServer, s)
	fmt.Printf("escuchando\n")
	grpcServer.Serve(listener)
	defer grpcServer.Stop()

}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func (s *NameNodeServer) SendLog(ctx context.Context, report *protos.Log) (*protos.Accept, error) {
	cont++
	accept := &protos.Accept{}
	if !contains(s.diccionario[report.NombreLibro+" "+report.CantidadPartes], report.Parte+" "+report.Ubicaciones) {
		s.diccionario[report.NombreLibro+" "+report.CantidadPartes] = append(s.diccionario[report.NombreLibro+" "+report.CantidadPartes], report.Parte+" "+report.Ubicaciones)
		fmt.Printf("Diccionario %v \n", s.diccionario)

	}
	i, _ := strconv.Atoi(report.CantidadPartes)
	if i == cont {
		makeLog(s)
		cont = 0
		s.diccionario = make(map[string][]string)
	}
	return accept, nil

}

func makeLog(s *NameNodeServer) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	for key := range s.diccionario {
		f.WriteString(key + "\n")
		for i := 0; i < len(s.diccionario[key]); i++ {
			f.WriteString(s.diccionario[key][i] + "\n")
		}
	}

}

func (s *NameNodeServer) RequestAdress(ctx context.Context, nombre *protos.Prop) (*protos.Adress, error) {
	adress := &protos.Adress{}
	//adress.Adress = make([]string)
	f, err := os.OpenFile("log.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	flag := false
	scanner := bufio.NewScanner(f)
	aux := 0
	aux2 := -1
	for scanner.Scan() {
		linea := scanner.Text()
		list := strings.Split((strings.Split(linea, "\n")[0]), " ")
		if flag {
			aux++
			ad := list[0] + " " + list[1]
			adress.Adress = append(adress.Adress, ad)
			if aux2 == aux {
				return adress, nil
			}

		}
		if list[0] == nombre.Node {
			i, _ := strconv.Atoi(list[1])
			aux2 = i
			flag = true
		}

		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return adress, nil

}
