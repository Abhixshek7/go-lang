package sendticket
import ("fmt"
	"time"
)

func Sendticket(userTicket uint, userName string,userEmail string){
	time.Sleep(10*time.Second)
	var ticket  = fmt.Sprintf("%v tickets for %v",userTicket,userName);
	fmt.Printf("Sending %v to email %v",ticket,userEmail)
}