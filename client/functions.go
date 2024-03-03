package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/mukesh/ticket_booking/proto"
)

func Book(client pb.TrainBookingClient, ticket *pb.TicketRequest) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("calling the book function")
	reponse, err := client.BookTicket(ctx, ticket)

	if err != nil {
		log.Printf("error in booking", err)
		return
	}

	log.Println("Seatno - ", reponse.Receipt.Seatno, "Coachno -", reponse.Receipt.Coachno.Section, " Status -", reponse.Status.Status)

}

func GetTicketDetail(client pb.TrainBookingClient, receipt *pb.TicketReceipt) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.GetReceipt(ctx, receipt)

	if err != nil {
		log.Printf("Error %v", err)
		return
	}

	log.Println("User from -", response.Userdetails.From, "User to -", response.Userdetails.To, "User Details - ", response.Userdetails.Users)
	log.Println("Seatno - ", response.Status.Receipt.Seatno, "Coachno -", response.Status.Receipt.Coachno.Section, " Status -", response.Status.Status.Status)
}

func CancelTicket(client pb.TrainBookingClient, receipt *pb.TicketReceipt) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := client.DeleteTicket(ctx, receipt)

	if err != nil {
		log.Printf("Error %v", err)
		return
	}

	fmt.Println("Your Ticket no", receipt.Seatno, "in this section", receipt.Coachno, "is ", response.Status)

}

func GetAllUsersData(client pb.TrainBookingClient, section *pb.TicketSection) {

	stream, err := client.GetAllUsers(context.Background(), section)
	if err != nil {
		log.Printf("error in reading stream data %v", err)
		return
	}
	log.Println("Receiving users Data...")
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}
		log.Println("User from ", response.Userdetails.From, "User to -", response.Userdetails.To, "User Details - ", response.Userdetails.Users)
		log.Println("Seatno - ", response.Status.Receipt.Seatno, "Coachno -", response.Status.Receipt.Coachno.Section, " Status -", response.Status.Status.Status)
	}
	log.Println("streaming completed")
}

func UpdateSeat(client pb.TrainBookingClient, receipt *pb.TicketReceipt) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reponse, err := client.ModifySeat(ctx, receipt)

	if err != nil {
		log.Printf("error in the changing seat", err)
		return
	}

	log.Println("Seatno - ", reponse.Receipt.Seatno, "Coachno -", reponse.Receipt.Coachno.Section, " Status -", reponse.Status.Status)

}
