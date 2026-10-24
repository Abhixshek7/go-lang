package main

import "fmt"

func main(){
	conferenceName := "Fucking Retard Group"
	const totalTickets = 69
	var availableTickets uint = 69

	bookings := [] string{}
	
	var userName string
	var userTicket uint

	for{
		fmt.Printf("Welcome to this %v of existentialists.!\n", conferenceName)
	fmt.Println("If you are here then you seriously need help. So, welcome to this shit show!")
	fmt.Printf("There are currently %v seats available out of %v tickets, be part of the greater good!\n",availableTickets,totalTickets);

	

	fmt.Println()
	fmt.Print("Enter your name you fucking retard : ")
	fmt.Scan(&userName)
	fmt.Print("How many seats do you want for you ass ? = ")
	fmt.Scan(&userTicket)

	fmt.Printf("Welcome to the shit show %v. You have purchased %v tickets.\n",userName,userTicket)
	bookings = append(bookings,userName);
	availableTickets -=userTicket
	fmt.Printf("%v tickets remaining(This is just for development purpose only)\n",availableTickets)
	fmt.Printf("%v \n",bookings)
	}
	// fmt.Println(conferenceName)
} 