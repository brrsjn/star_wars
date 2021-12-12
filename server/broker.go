package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	pb "star_wars/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50050"

	address1 = "localhost:50051"
	address2 = "localhost:50052"
	address3 = "localhost:50053"
)

type Broker struct {
	pb.UnimplementedInformerToBrokerServer
	servers    [3]*pb.Servidor
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
	pb.RegisterInformerToBrokerServer(s, &Broker{
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

func (broker *Broker) ConnectToServer(ctx context.Context, req *pb.Instruct) (*pb.Servidor, error) {
	if broker.conectedSV > 0 {
		req := broker.servers[rand.Intn(2)]
		return req, nil
	} else {
		return &pb.Servidor{Addres: "nada"}, nil
	}

}

func RandomServer() string {
	adressServer := "empty"
	return adressServer
}
