package main

import (
	"strings"
	"unicode"
)
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