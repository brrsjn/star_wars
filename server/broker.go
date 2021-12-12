package main

import (
	"log"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// hay que hacer un proto con los servidores y suscribirlo aqui
	//nameNode := proto.NewServerNodeServiceClient(conn)

	//s := grpc.NewServer()

	//algo := protos.RegisterInformerToBrokerServer(s, )
}

func RandomServer() string {
	adressServer := "empty"
	return adressServer
}
