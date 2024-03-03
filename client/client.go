package main

import (
	"fmt"
	"log"

	pb "github.com/mukesh/ticket_booking/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serveraddr = "127.0.0.1:9090"
)

func main() {

	conn, err := grpc.Dial(serveraddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("cannot make connection to the grpc server", err)
	}

	client := pb.NewTrainBookingClient(conn)

	fmt.Print("1 - for booking the ticket \n2 - for Getting the Details \n3 - for - Viewing All the Users \n" +
		"4 - for Modifying the User \n5 - for Deleting \n6 - for exiting the Application \n")

	for {

		var number int64
		fmt.Println("Select the displayed option")
		_, err := fmt.Scan(&number)
		if err != nil {
			log.Println("enter correct input data")
		}
		fmt.Println("you have selected this number", number)
		switch number {
		case 1:
			fmt.Println("enter the ticket details to book the data")
			ticket, err := CreateTicketRequest()
			if err != nil {
				log.Println("error ---", err)
				return
			}
			Book(client, ticket)

		case 2:
			fmt.Println("enter the ticket details to Get the data")
			receipt, err := CreateTicketReceipt()
			if err != nil {
				log.Println("error ---", err)
				return
			}
			GetTicketDetail(client, receipt)
		case 3:
			fmt.Println("enter the Section details to Get ALL section data")
			platform, err := CreateTicketSection()
			if err != nil {
				log.Println("error ---", err)
				return
			}
			GetAllUsersData(client, platform)
		case 4:
			fmt.Println("enter the Receipt details to Modify seat")

			receipt, err := CreateTicketReceipt()
			if err != nil {
				log.Println("error ---", err)
				return
			}

			UpdateSeat(client, receipt)
		case 5:
			fmt.Println("enter the Receipt details to delete the seat")
			receipt, err := CreateTicketReceipt()
			if err != nil {
				log.Println("error ---", err)
				return
			}
			CancelTicket(client, receipt)
		case 6:
			log.Println("exiting ....")
			return
		}

	}

	defer conn.Close()
}
