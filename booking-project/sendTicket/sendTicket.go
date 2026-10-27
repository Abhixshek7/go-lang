	package sendticket
	import ("fmt"
		"time"
		"sync"
	)

	func Sendticket(userTicket uint, userName string,userEmail string,wg *sync.WaitGroup){
		defer wg.Done()
		time.Sleep(5*time.Second)
		var ticket  = fmt.Sprintf("%v tickets for %v",userTicket,userName);
		fmt.Println("###########################################\n")
		fmt.Printf("Sending %v to email %v",ticket,userEmail)
		fmt.Println("\n###########################################\n")
	}