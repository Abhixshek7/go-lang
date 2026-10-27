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
type UserData struct{
	UserName string
	UserEmail string
	UserTicket uint
}

var bookings = make([]UserData,0)

func main(){
	

	// bookings := [] string{}

	
	for{
	greet.GreetMsg(ConferenceName,availableTickets,totalTickets)
	userName,userEmail,userTicket := gatherUserData.GatherUserData(availableTickets);
	availableTickets  = booking.Booking(userTicket,availableTickets,userName,userEmail)
	fmt.Println()
	
	newBooking := UserData{
		UserName :userName,
		UserEmail : userEmail,
		UserTicket: userTicket,
	}

	bookings = append(bookings,newBooking)
	if availableTickets ==0{
		fmt.Println("The bookings are over. so please fuck off!!")
		fmt.Println("\n--- Booking List (For Development Purpose Only) ---")
		for _,booking :=range(bookings){
			fmt.Printf("Name : %v | Email : %v | Tickets : %v",
			booking.UserName,
			booking.UserEmail,
			booking.UserTicket,);
			fmt.Println()
		}
		break;
	}
	
	}
	// fmt.Println(conferenceName)
} 



