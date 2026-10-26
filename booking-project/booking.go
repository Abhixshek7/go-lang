package main
import "fmt"
func booking(userTicket uint, availableTickets uint, userName string,userEmail string) uint{
	if(userTicket>availableTickets){
		fmt.Print("Seat limit exceeds, try again!!")
		return availableTickets
	}

	fmt.Printf("Welcome to the shit show %v. You have purchased %v tickets.Confirmation email sent to the following account = %v \n",userName,userTicket,userEmail)
	// bookings = append(bookings,userName);
	availableTickets -=userTicket
	fmt.Printf("%v tickets remaining(This is just for development purpose only)\n",availableTickets)
	

	// if(availableTickets==0){
	// 	fmt.Print("The tickets are sold out so please fuck off")
	// }
	return availableTickets
}