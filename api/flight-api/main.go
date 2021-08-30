package main

import (
	"Assignment-micro-service/api/flight-api/handlers"
	"Assignment-micro-service/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":2223", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	flightClient := pb.NewFlightsClient(conn)
	h := handlers.NewFlightHandler(flightClient)
	g := gin.Default()
	gr := g.Group("/flight")
	gr.POST("/create", h.CreateFlight)
	gr.POST("/update", h.UpdateFlight)
	gr.POST("/search", h.SearchFly)
	http.ListenAndServe(":4444", g)
}
