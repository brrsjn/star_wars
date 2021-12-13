package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	broker "star_wars/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50050"

	address1 = "localhost:50051"
	address2 = "localhost:50052"
	address3 = "localhost:50053"
)

type BrokerObj struct {
	broker.UnimplementedBrokerServer
	servers    [3]*broker.Servidor
	conectedSV int32
}

func main() {
	// Conectar con los servidores de Fulcrum

	/*lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	*/
	// hay que hacer un proto con los servidores y suscribirlo aqui
	//nameNode := proto.NewServerNodeServiceClient(conn)
	//algo := protos.RegisterInformerToBrokerServer(s, )

	//Iniciar como servidor
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	broker.RegisterBrokerServer(s, &BrokerObj{
		conectedSV: 0,
	})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func ConnectToFulcrum(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
}

func (self *BrokerObj) ConnectToServer(ctx context.Context, req *broker.Instruct) (*broker.Servidor, error) {
	if self.conectedSV == 0 {
		fmt.Println("-Mos Eisley: No hay servidores disponibles...")
		return &broker.Servidor{Addres: "empty"}, nil

	} else {
		fmt.Println("-Mos Eisley: Comunicando...")
		req := self.servers[rand.Intn(2)]
		return req, nil
	}

}

func RandomServer() string {
	adressServer := "empty"
	return adressServer
}
