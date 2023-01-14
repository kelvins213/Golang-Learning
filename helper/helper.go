package helper

import (
	"fmt"
	"booking-app/getters"
);


var Bookings []string;
var RemainingTickets uint = 50;
const ConferenceTickets uint = 50;

func BuyTickets(){
	RemainingTickets = RemainingTickets - getters.UserTickets;
	Bookings = append(Bookings, getters.UserName);
	fmt.Println("Your tickets has just been granted");
}

