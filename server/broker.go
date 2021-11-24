package main

import (
	"context"
	"log"
	"net"
	"star_wars/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address1 = "localhost:50061"
	address2 = "localhost:50062"
	address3 = "localhost:50063"
	port     = ":50051"
)

type InformanteBrokerServer struct {
	pb.UnimplementedInformanteBrokerServiceServer
	cliente1 pb.BrokerFulcrumServiceClient
	cliente2 pb.BrokerFulcrumServiceClient
	cliente3 pb.BrokerFulcrumServiceClient
}

func (s *InformanteBrokerServer) AddCityToBroker(ctx context.Context, in *pb.AddCityToBrokerRequest) (*pb.AddCityToBrokerReply, error) {
	//Codigo para guardar la ciudad en archivo
	return &pb.AddCityToBrokerReply{State: true}, nil
}

func (s *InformanteBrokerServer) DeleteCityToBroker(ctx context.Context, in *pb.DeleteCityToBrokerRequest) (*pb.DeleteCityToBrokerReply, error) {
	//Codigo para eliminar una ciudad
	return &pb.DeleteCityToBrokerReply{State: true}, nil
}

func (s *InformanteBrokerServer) UpdateNameToBroker(ctx context.Context, in *pb.UpdateNameToBrokerRequest) (*pb.UpdateNameToBrokerReply, error) {
	//Codigo para eliminar una ciudad
	return &pb.UpdateNameToBrokerReply{State: true}, nil
}

func (s *InformanteBrokerServer) UpdateNumberToBroker(ctx context.Context, in *pb.UpdateNumberToBrokerRequest) (*pb.UpdateNumberToBrokerReply, error) {
	//Codigo para eliminar una ciudad
	return &pb.UpdateNumberToBrokerReply{State: true}, nil
}

func main() {

	//Connection to te fulkrum server

	// Set up a connection to the server.
	conn1, err := grpc.Dial(address1, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn1.Close()
	c1 := pb.NewBrokerFulcrumServiceClient(conn1)

	// Set up a connection to the server.
	conn2, err := grpc.Dial(address2, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()
	c2 := pb.NewBrokerFulcrumServiceClient(conn2)

	// Set up a connection to the server.
	conn3, err := grpc.Dial(address3, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn3.Close()
	c3 := pb.NewBrokerFulcrumServiceClient(conn3)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterInformanteBrokerServiceServer(s, &InformanteBrokerServer{
		cliente1: c1,
		cliente2: c2,
		cliente3: c3,
	})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
