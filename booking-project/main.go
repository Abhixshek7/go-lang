package main

import (
	"fmt"

)

var conferenceName = "Fucking Retard Group"
const totalTickets = 69
var availableTickets uint = 69

func main(){
	

	// bookings := [] string{}

	
	for{
	greetMsg()
	userName,userEmail,userTicket := gatherUserData();
	availableTickets  = booking(userTicket,availableTickets,userName,userEmail)
	fmt.Println()
	
	if availableTickets ==0{
		fmt.Println("The bookings are over. so please fuck off!!")
		break;
	}
	
	}
	// fmt.Println(conferenceName)
} 



