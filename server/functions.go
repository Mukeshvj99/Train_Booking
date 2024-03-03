package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/mukesh/ticket_booking/proto"
)

func (s *server) BookTicket(ctx context.Context, ticket *pb.TicketRequest) (*pb.TicketResponse, error) {

	log.Println("started booking...")

	if coach.Alen < len(coach.A) {
		log.Println("enteredd the coach A")
		seatno := coach.findSeat("A", ticket, -1)
		fmt.Println("seat", seatno)
		if seatno >= 0 {

			log.Println("Seat confirmed")
			return SeatConfirmation(seatno, "A", "successfully booked")
		}

	} else if coach.Blen < len(coach.B) {
		seatno := coach.findSeat("B", ticket, -1)
		if seatno >= 0 {
			log.Printf("Seat confirmed %v", ticket.Users.Firstname)
			return SeatConfirmation(seatno, "B", "successfully booked")
		}

	}

	return &pb.TicketResponse{}, fmt.Errorf("Train is fully booked")
}

func (c *section) findSeat(platform string, ticket *pb.TicketRequest, currseat int) int {

	mu.Lock()
	var seat int = -1
	if platform == "A" {

		for ind, val := range c.A {

			if val == nil {

				if currseat != -1 {
					c.A[currseat] = nil
					fmt.Println("after updation", c.A[currseat])
					c.Alen--
				}
				c.A[ind] = ticket
				c.Alen++
				seat = ind
				break
			}
		}

	} else if platform == "B" {
		for ind, val := range c.B {

			if val == nil {

				if currseat != -1 {
					c.B[currseat] = nil
					c.Blen--
				}
				c.B[ind] = ticket
				c.Blen++
				seat = ind
				break
			}
		}
	}

	mu.Unlock()

	return seat

}

func SeatConfirmation(seatid int, coach, status string) (*pb.TicketResponse, error) {

	return &pb.TicketResponse{

		Receipt: &pb.TicketReceipt{
			Seatno: &pb.TicketNumber{
				Ticket: int32(seatid) + 1,
			},
			Coachno: &pb.TicketSection{
				Section: coach,
			},
		},
		Status: &pb.TicketStatus{
			Status: status,
		},
	}, nil
}

func (s *server) GetReceipt(ctx context.Context, receipt *pb.TicketReceipt) (*pb.TicketDetails, error) {

	ticketno := receipt.Seatno.Ticket
	section := receipt.Coachno.Section

	presence, err := coach.VerifySeat(int(ticketno), section)
	if err != nil {
		return nil, err
	}

	if presence {
		if section == "A" {
			return &pb.TicketDetails{
				Userdetails: coach.A[ticketno-1],
				Status: &pb.TicketResponse{
					Receipt: &pb.TicketReceipt{
						Seatno: &pb.TicketNumber{
							Ticket: ticketno,
						},
						Coachno: &pb.TicketSection{
							Section: "A",
						},
					},
					Status: &pb.TicketStatus{
						Status: "Successfully booked",
					},
				},
			}, nil

		} else {
			return &pb.TicketDetails{
				Userdetails: coach.B[ticketno-1],
				Status: &pb.TicketResponse{
					Receipt: &pb.TicketReceipt{
						Seatno: &pb.TicketNumber{
							Ticket: ticketno,
						},
						Coachno: &pb.TicketSection{
							Section: "B",
						},
					},
					Status: &pb.TicketStatus{
						Status: "Successfully booked",
					},
				},
			}, nil
		}
	}
	return nil, nil
}

func (s *section) VerifySeat(ticket int, section string) (bool, error) {

	if section == "A" {

		if ticket > len(s.A) || s.A[ticket-1] == nil {
			return false, fmt.Errorf("ticket doesn't exist")
		}

	} else if section == "B" {
		if ticket > len(s.B) || s.B[ticket-1] == nil {
			return false, fmt.Errorf("ticket doesn't exist")
		}

	}

	if (section == "A" || section == "B") && (ticket <= len(s.A) || ticket <= len(s.B)) {
		return true, nil
	}

	return false, fmt.Errorf("Section Doesn't exist")
}

func (s *server) DeleteTicket(ctx context.Context, receipt *pb.TicketReceipt) (*pb.TicketStatus, error) {

	ticketno := receipt.Seatno.Ticket
	section := receipt.Coachno.Section

	_, err := coach.VerifySeat(int(ticketno), section)
	if err != nil {
		return nil, err
	}

	status := coach.DeleteUser(int(ticketno), section)

	return &pb.TicketStatus{
		Status: status,
	}, nil

}

func (c *section) DeleteUser(ticket int, coach string) string {

	mu.Lock()
	status := "Not Cancelled"
	if coach == "A" {
		c.A[ticket-1] = nil
		c.Alen--
		status = "Successfully Cancelled"
	} else if coach == "B" {
		c.B[ticket-1] = nil
		c.Blen--
		status = "Successfully Cancelled"
	}
	mu.Unlock()

	return status

}

func (s *server) GetAllUsers(sectionno *pb.TicketSection, stream pb.TrainBooking_GetAllUsersServer) error {

	platform := sectionno.Section

	if platform != "A" && platform != "B" {
		return fmt.Errorf("Invalid Section")
	}
	log.Println("streaming started for ", platform)

	for _, res := range coach.Read(platform) {

		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) ModifySeat(ctx context.Context, receipt *pb.TicketReceipt) (*pb.TicketResponse, error) {

	log.Println("started modifying the seat...")

	ticketno := receipt.Seatno.Ticket
	section := receipt.Coachno.Section

	_, err := coach.VerifySeat(int(ticketno), section)
	if err != nil {
		return nil, err
	}

	if coach.Alen < len(coach.A) {
		log.Println("enteredd the coach A")
		ticket := coach.A[ticketno-1]
		seatno := coach.findSeat("A", ticket, int(ticketno-1))
		fmt.Println("seat", seatno)
		if seatno >= 0 {

			log.Println("Seat confirmed")
			return SeatConfirmation(seatno, "A", "successfully booked")
		}

	} else if coach.Blen < len(coach.B) {
		ticket := coach.B[ticketno-1]

		seatno := coach.findSeat("B", ticket, int(ticketno-1))
		if seatno >= 0 {
			log.Printf("Seat confirmed %v", ticket.Users.Firstname)
			return SeatConfirmation(seatno, "B", "successfully booked")
		}

	}

	return &pb.TicketResponse{}, fmt.Errorf("Train is fully booked")
}
func (c *section) Read(coach string) []*pb.TicketDetails {

	UserDetails := make([]*pb.TicketDetails, 0)
	var arr []*pb.TicketRequest
	if coach == "A" {
		arr = c.A[:]
	} else if coach == "B" {
		arr = c.B[:]
	}
	for ind, val := range arr {

		if val != nil {
			response := &pb.TicketDetails{
				Userdetails: val,
				Status: &pb.TicketResponse{
					Receipt: &pb.TicketReceipt{
						Seatno: &pb.TicketNumber{
							Ticket: int32(ind + 1),
						},
						Coachno: &pb.TicketSection{
							Section: "A",
						},
					},
					Status: &pb.TicketStatus{
						Status: "Successfully booked",
					},
				},
			}

			UserDetails = append(UserDetails, response)
		}
	}
	return UserDetails
}
