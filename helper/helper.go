package helper

import (
	"fmt"
	"strconv"
	"booking-app/getters"
);

/*
The data type Map is like an array which allows us to store data in each 
index like a json.
*/

var Bookings []string;
var BookingsNames []string;
var RemainingTickets uint = 50;
const ConferenceTickets uint = 50;

func BuyTickets(){

	var BookingsNamesMap = make(map[string]string,0); 


	RemainingTickets = RemainingTickets - getters.UserTickets;
	Bookings = append(Bookings, getters.UserName);

	str := strconv.FormatUint(uint64(getters.UserTickets), 10);

	BookingsNamesMap["Name"] = getters.UserName;
	BookingsNamesMap["Space"] = " ";
	BookingsNamesMap["Tickets"] = str;

	BookingsNames = append(BookingsNames, BookingsNamesMap["Name"] + BookingsNamesMap["Space"] + BookingsNamesMap["Tickets"] + BookingsNamesMap["Space"] + "Tickets");
	fmt.Println("Your tickets has just been granted");
}

