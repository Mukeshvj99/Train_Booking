package main

import (
	"fmt"
	"log"
	"strings"

	pb "github.com/mukesh/ticket_booking/proto"
)

func getFrom() (string, error) {
	var from string
	fmt.Println("enter from:")
	_, err := fmt.Scan(&from)
	if err != nil {
		log.Printf("error in reading the from data from  users %v", err)
		return "", err
	}

	return from, nil
}

func getTo() (string, error) {
	var to string
	fmt.Println("enter to:")
	_, err := fmt.Scan(&to)
	if err != nil {
		log.Printf("error in reading the to data from  users %v", err)
		return "", err
	}

	return to, nil
}

func getFirstname() (string, error) {
	var firstname string
	fmt.Println("enter firstname:")
	_, err := fmt.Scan(&firstname)
	if err != nil {
		log.Printf("error in reading the firstname data from  users %v", err)
		return "", err
	}

	return firstname, nil
}

func getLastname() (string, error) {
	var lastname string
	fmt.Println("enter lastname:")
	_, err := fmt.Scan(&lastname)
	if err != nil {
		log.Printf("error in reading the lastname data from  users %v", err)
		return "", err
	}

	return lastname, nil
}

func getEmail() (string, error) {
	var email string
	fmt.Println("enter Email:")
	_, err := fmt.Scan(&email)
	if err != nil {
		log.Printf("error in reading the email data from  users %v", err)
		return "", err
	}

	return email, nil
}

func getPrice() (float32, error) {
	var price float32
	fmt.Println("enter price")
	_, err := fmt.Scan(&price)
	if err != nil {
		log.Printf("error in reading the price data from  users %v", err)
		return 0, err
	}

	return price, nil
}

func getTicket() (int, error) {
	var ticket int
	fmt.Println("enter ticket no")
	_, err := fmt.Scan(&ticket)
	if err != nil {
		log.Println("error in reading the ticket ", err)
		return -1, nil
	}

	return ticket, nil
}

func getCoach() (string, error) {
	var coachno string
	fmt.Println("enter coach no")
	_, err1 := fmt.Scan(&coachno)
	if err1 != nil {
		log.Println("error in reading the coachno ", err1)
		return "", nil
	}

	return strings.ToUpper(coachno), nil
}

func CreateTicketRequest() (*pb.TicketRequest, error) {

	from, err := getFrom()
	if err != nil {
		return nil, err
	}
	to, err := getTo()
	if err != nil {
		return nil, err
	}

	price, err := getPrice()
	if err != nil {
		return nil, err
	}

	user, err := CreateUser()
	if err != nil {
		return nil, err
	}
	ticket := &pb.TicketRequest{
		From:  from,
		To:    to,
		Users: user,
		Price: price,
	}

	return ticket, nil
}
func CreateUser() (*pb.User, error) {

	firstname, err := getFirstname()
	if err != nil {
		return nil, err
	}

	lastname, err := getLastname()
	if err != nil {
		return nil, err
	}

	email, err := getEmail()
	if err != nil {
		return nil, err
	}

	user := &pb.User{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
	}

	return user, nil
}
func CreateTicketReceipt() (*pb.TicketReceipt, error) {
	ticket, err := getTicket()
	if err != nil {
		return nil, err
	}
	coach, err := getCoach()
	if err != nil {
		return nil, err
	}

	receipt := pb.TicketReceipt{
		Seatno: &pb.TicketNumber{
			Ticket: int32(ticket),
		},
		Coachno: &pb.TicketSection{
			Section: coach,
		},
	}

	return &receipt, nil
}

func CreateTicketSection() (*pb.TicketSection, error) {
	coach, err := getCoach()
	if err != nil {
		return nil, err
	}

	platform := *&pb.TicketSection{
		Section: coach,
	}

	return &platform, nil
}
