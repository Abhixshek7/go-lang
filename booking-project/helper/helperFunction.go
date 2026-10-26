package helper

import (
	"strings"
	"unicode"
)
func ValidateName(userName string) bool{
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


func ValidateEmail(userEmail string)bool{
	return strings.Contains(userEmail,"@") && strings.Contains(userEmail,".")
}

func ValidateTickeCount(userTicket uint,availableTickets uint) bool{
	return userTicket>0 && userTicket<=availableTickets ;
}