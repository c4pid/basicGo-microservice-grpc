package main

import (
	"Assignment-micro-service/gprc-customers/handlers"
	"Assignment-micro-service/gprc-customers/repositories"
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
	listen, err := net.Listen("tcp", ":2222")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	customerRepository, err := repositories.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := handlers.NewCustomerHandler(customerRepository)
	if err != nil {
		panic(err)
	}

	reflection.Register(s)
	pb.RegisterCustomerServiceServer(s, h)
	fmt.Println("Listen on port 2222")
	s.Serve(listen)
}
