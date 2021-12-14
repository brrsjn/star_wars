package main

import (
	"context"
	"log"
	"star_wars/pb"
	"time"

	"google.golang.org/grpc"
)

const (
	address    = "localhost:50061"
	defaultBot = true
)

func testAddCity() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFulcrumClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddCity(ctx, &pb.City{Name: "JuanCity2", Planet: "Tatooine", Survivors: 3})
	if err != nil {
		log.Fatalf("could not greet: %v", err)

	}
	log.Printf("Ciudad %s creada", r.Name)
}

func testUpdateName() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFulcrumClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.UpdateName(ctx, &pb.CityNewName{City: "JuanCity2", Planet: "Tatooine", NewName: "SolCity2"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)

	}
	log.Printf("Ciudad %s Actualizada", r.Name)
}

func testUpdateNumber() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFulcrumClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.UpdateNumber(ctx, &pb.CityNewNumber{City: "JuanCity2", Planet: "Tatooine", Survivors: 5})
	if err != nil {
		log.Fatalf("could not greet: %v", err)

	}
	log.Printf("Ciudad %s Actualizada su numero de survivors", r.Name)
}

func main() {
	testAddCity()
	//testUpdateName()
	//testUpdateNumber()
}
