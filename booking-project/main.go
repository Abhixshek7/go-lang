package main

import (
	"fmt"
	"booking-project/booking"
	"booking-project/gatherUserData"
	"booking-project/greet"
	
	

)

var ConferenceName = "Fucking Retard Group"
const totalTickets = 69
var availableTickets uint = 69

func main(){
	

	// bookings := [] string{}

	
	for{
	greet.GreetMsg(ConferenceName,availableTickets,totalTickets)
	userName,userEmail,userTicket := gatherUserData.GatherUserData(availableTickets);
	availableTickets  = booking.Booking(userTicket,availableTickets,userName,userEmail)
	fmt.Println()
	
	if availableTickets ==0{
		fmt.Println("The bookings are over. so please fuck off!!")
		break;
	}
	
	}
	// fmt.Println(conferenceName)
} 



