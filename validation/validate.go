package validation

import (
	"fmt"
	"booking-app/helper"
	"booking-app/main"
	"booking-app/geters"
);

var RemainingTickets uint = 50;

func ValidateUserInputs(){
	var userNameIsValid = len(UserName) <= 2;
	var ticketsAreValide = UserTickets > RemainingTickets;

	if (!userNameIsValid && !ticketsAreValide) {
		helper.BuyTickets();
	} else {
		if (userNameIsValid && ticketsAreValide) {
			fmt.Printf("\nSorry, your name has less than 2 letters \n");
			fmt.Printf("\nSorry, we only have %d tickets left \n", RemainingTickets);
			fmt.Printf("Try again \n\n");
			main.ShowMenu();
		} else {
			if (userNameIsValid) {
				fmt.Printf("\nSorry, your name has less than 2 letters \n\n");
				main.ShowMenu();
			} else {
				fmt.Printf("\nSorry, we only have %d tickets left \n", RemainingTickets);
				fmt.Printf("Please, try again \n\n");
				main.ShowMenu();
			}
		}
	}
}
