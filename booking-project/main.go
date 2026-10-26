package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main(){
	conferenceName := "Fucking Retard Group"
	const totalTickets = 69
	var availableTickets uint = 69

	// bookings := [] string{}

	
	for{
	greetMsg(conferenceName,availableTickets,totalTickets)
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

func greetMsg(confName string,availableTickets uint, totalTickets int){
	fmt.Printf("Welcome to this %v of existentialists.!\n", confName)
	fmt.Println("If you are here then you seriously need help. So, welcome to this shit show!")
	fmt.Printf("There are currently %v seats available out of %v tickets, be part of the greater good!\n",availableTickets,totalTickets);
	fmt.Println();
}

func gatherUserData()(string, string,uint){
	

	var userName string
	var userTicket uint
	var userEmail string

	for{
		fmt.Print("Enter your name you fucking retard : ")
		fmt.Scan(&userName)
		if validateName(userName){
			break
		}

		fmt.Println("Enter a valid name you fucking retard. A name cannot be empty| cannot have numbers, unless you are a nigga.")
	}

	for{

		fmt.Print("Enter your email address which u use to visit the hub : ")
		fmt.Scan(&userEmail)

		if validEmail(userEmail){
			break
		}
		fmt.Println("Dont fucking lie man, you cant get the hub mail wrong!!!!")
	}

	for{
		
		fmt.Print("How many seats do you want for you ass ? = ")
		fmt.Scan(&userTicket)
		if validTickeCount(userTicket){
			break
		}
		fmt.Println("Enter a valid number u retard!!")
	}
	


	return userName,userEmail,userTicket
}

func validateName(userName string) bool{
	if userName ==" " || len(userName)<2{
		return false
	}

	for _,char :=range(userName){
		if !unicode.IsLetter(char) && char != ' '{
			return false
		}
	}
	return true;
}


func validEmail(userEmail string)bool{
	return strings.Contains(userEmail,"@") && strings.Contains(userEmail,".")
}

func validTickeCount(userTicket uint) bool{
	return userTicket>0;
}
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