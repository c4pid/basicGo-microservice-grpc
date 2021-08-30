package main

import (
	"Assignment-micro-service/grpc-flights/handlers"
	"Assignment-micro-service/grpc-flights/repositories"
	"Assignment-micro-service/helper"
	"Assignment-micro-service/pb"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := helper.AutoBindConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", ":2223")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	flightRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewFlightHandler(flightRepository)
	if err != nil {
		panic(err)
	}
	reflection.Register(s)
	pb.RegisterFlightsServer(s, h)
	fmt.Println("Listen on port 2223")
	s.Serve(listen)
}
