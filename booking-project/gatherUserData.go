package main
import "fmt"
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