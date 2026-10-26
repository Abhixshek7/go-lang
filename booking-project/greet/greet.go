package greet

import "fmt"
func GreetMsg(ConferenceName string,availableTickets uint, totalTickets int){
	fmt.Printf("Welcome to this %v of existentialists.!\n", ConferenceName)
	fmt.Println("If you are here then you seriously need help. So, welcome to this shit show!")
	fmt.Printf("There are currently %v seats available out of %v tickets, be part of the greater good!\n",availableTickets,totalTickets);
	fmt.Println();
}