package main

import (
	"Assignment-micro-service/api/customer-api/handlers"
	"Assignment-micro-service/pb"

	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":2222", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	customerClient := pb.NewCustomerServiceClient(conn)

	h := handlers.NewCustomerHandler(customerClient)
	g := gin.Default()
	gr := g.Group("/customer")
	gr.POST("/create", h.CreateCustomer)
	gr.POST("/update", h.UpdateCustomer)
	gr.POST("/changepass", h.ChangePassword)
	http.ListenAndServe(":3333", g)
}
