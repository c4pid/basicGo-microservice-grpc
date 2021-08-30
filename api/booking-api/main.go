package main

import (
	"Assignment-micro-service/api/booking-api/handlers"
	"Assignment-micro-service/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":2224", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	bookingClient := pb.NewBookingsClient(conn)
	h := handlers.NewBookingClient(bookingClient)
	g := gin.Default()
	gr := g.Group("/booking")
	gr.POST("/create", h.Booking)
	gr.POST("/view", h.ViewBooking)
	gr.POST("/cancel", h.CancelBooking)
	http.ListenAndServe(":5555", g)
}
