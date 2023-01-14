package menu

import (
	"fmt";
	"booking-app/listing"
	"booking-app/getters"
	"booking-app/validation"
);

func ShowMenu(){
	var wish int;

	
	//bookings := [49] string{};
	//var bookings = [50]string{"Nana", "Nicole", "Peter"};

	for (wish != 5 && validation.RemainingTickets!= 0) {
		fmt.Println("================> MENU <================");
		fmt.Printf("1 - Buy tickets \n2 - List Buyers \n3 - List Buyers By Tickets \n4 - List Remaining Tickets \n5 - Exit System \n: ");
		fmt.Scan(&wish)

		switch wish {
			case 1:
				userInteration();
			case 2:
				listing.ListBuyers();
			case 3:
				listing.ListBuyersByTickets();
			case 4:
				listing.ListRemainingTickets();
			default:
				endMenu();
		}
	}
}

func endMenu(){
	fmt.Println("Thanks, goodbye!");
}


func userInteration(){
	getters.GetUserName();
	getters.GetUserTickets();
	validation.ValidateUserInputs();
}