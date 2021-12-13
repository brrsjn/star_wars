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
	port = ":50061"
)

type BrokerFulcrumServer struct {
	pb.UnimplementedBrokerFulcrumServiceServer
}

func (s *BrokerFulcrumServer) AddCityToFulcrum(ctx context.Context, in *pb.AddCityToFulcrumRequest) (*pb.AddCityToFulcrumReply, error) {
	//Codigo para guardar la ciudad en archivo
	return &pb.AddCityToFulcrumReply{State: true}, nil
}

func (s *BrokerFulcrumServer) DeleteCityToFulcrum(ctx context.Context, in *pb.DeleteCityToFulcrumRequest) (*pb.DeleteCityToFulcrumReply, error) {
	//Codigo para eliminar una ciudad del archivo
	return &pb.DeleteCityToFulcrumReply{State: true}, nil
}

func (s *BrokerFulcrumServer) UpdateNameToFulcrum(ctx context.Context, in *pb.UpdateNameToFulcrumRequest) (*pb.UpdateNameToFulcrumReply, error) {
	//Codigo para actualizar una ciudad
	return &pb.UpdateNameToFulcrumReply{State: true}, nil
}

func (s *BrokerFulcrumServer) UpdateNumberToFulcrum(ctx context.Context, in *pb.UpdateNumberToFulcrumRequest) (*pb.UpdateNumberToFulcrumReply, error) {
	//Codigo cambiar el numero de la ciudad
	return &pb.UpdateNumberToFulcrumReply{State: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBrokerFulcrumServiceServer(s, &BrokerFulcrumServer{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
