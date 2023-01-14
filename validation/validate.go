package validation

import (
	"booking-app/getters"
	"booking-app/helper"
	//"booking-app/menu"
	"fmt"
);

var RemainingTickets uint = 50;

func ValidateUserInputs(){
	var userNameIsValid = len(getters.UserName) <= 2;
	var ticketsAreValide = getters.UserTickets > RemainingTickets;

	if (!userNameIsValid && !ticketsAreValide) {
		helper.BuyTickets();
	} else {
		if (userNameIsValid && ticketsAreValide) {
			fmt.Printf("\nSorry, your name has less than 2 letters \n");
			fmt.Printf("\nSorry, we only have %d tickets left \n", RemainingTickets);
			fmt.Printf("Try again \n\n");
		} else {
			if (userNameIsValid) {
				fmt.Printf("\nSorry, your name has less than 2 letters \n\n");
			} else {
				fmt.Printf("\nSorry, we only have %d tickets left \n", RemainingTickets);
				fmt.Printf("Please, try again \n\n");
			}
		}
	}
}
