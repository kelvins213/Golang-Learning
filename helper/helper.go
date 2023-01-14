package helper

import (
	"fmt"
	"booking-app/geters"
);


var bookings []string;
var remainingTickets uint = 50;

func BuyTickets(){
	remainingTickets = remainingTickets - UserTickets;
	bookings = append(bookings, geters.UserName);
	fmt.Println("Your tickets has just been granted");
}

