package listing;


import (
	"fmt"
	"booking-app/helper";
);

func ListBuyers(){
	var arrayLength = len(helper.Bookings);
	if (arrayLength > 0) {
		for index, buyer := range helper.Bookings {
			//index => the position where buyer is arrenged in the bookings array
			//buyer => its the element
			fmt.Printf("%dº buyer: [%s] \n", index + 1, buyer);
		}
		/*
		for _, buyer := range bookings{} the underline works to ignore the use of a expected variable
		*/
	} else {
		fmt.Println("We don't have buyers yet!");
	}
}

func ListRemainingTickets(){
	fmt.Printf("\n We still have %d tickets avaliable of %d tickets \n", helper.RemainingTickets, helper.ConferenceTickets);
}