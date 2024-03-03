package main

import (
	"log"
	"net"
	"sync"

	pb "github.com/mukesh/ticket_booking/proto"
	"google.golang.org/grpc"
)

var (
	addr = "127.0.0.1:9090"
	mu   sync.Mutex
)

type server struct {
	pb.TrainBookingServer
}

type section struct {
	A    [5]*pb.TicketRequest
	B    [3]*pb.TicketRequest
	Alen int
	Blen int
}

var coach *section

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("cannot listen on the given address", err)
	}

	log.Println("server started")

	grpcserver := grpc.NewServer()
	log.Println("Initializing the train sections... ")
	var sectionA [5]*pb.TicketRequest
	var sectionB [3]*pb.TicketRequest
	coach = &section{
		A:    sectionA,
		B:    sectionB,
		Alen: 0,
		Blen: 0,
	}

	pb.RegisterTrainBookingServer(grpcserver, &server{})
	log.Println("server initalized")
	if err = grpcserver.Serve(lis); err != nil {
		log.Fatalf("cannot serve the Request", err)
	}

	log.Println("server stopped....")
}
