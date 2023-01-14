package main

import "fmt";

func getUserName() {
	fmt.Printf("Please, tell me your name: ");
	fmt.Scanf("%v", &userName); 
}

func getUserTickets() {
	fmt.Printf("\nHow many tickets do you wish to buy? Answer: ");
	fmt.Scanf("%d", &userTickets); //o & serve para o compilador também ler o endereço de memória da variável userTickets
}

func validateUserInputs(){
	var userNameIsValid = len(userName) <= 2;
	var ticketsAreValide = userTickets > remainingTickets;

	if (!userNameIsValid && !ticketsAreValide) {
		buyTickets();
	} else {
		if (userNameIsValid && ticketsAreValide) {
			fmt.Printf("\nSorry, your name has less than 2 letters \n");
			fmt.Printf("\nSorry, we only have %d tickets left \n", remainingTickets);
			fmt.Printf("Try again \n\n");
			showMenu();
		} else {
			if (userNameIsValid) {
				fmt.Printf("\nSorry, your name has less than 2 letters \n\n");
				showMenu();
			} else {
				fmt.Printf("\nSorry, we only have %d tickets left \n", remainingTickets);
				fmt.Printf("Please, try again \n\n");
				showMenu();
			}
		}
	}
}

func buyTickets(){
	remainingTickets = remainingTickets - userTickets;
	bookings = append(bookings, userName);
	fmt.Println("Your tickets has just been granted");
}

func listBuyers(){
	var arrayLength = len(bookings);
	if (arrayLength > 0) {
		for index, buyer := range bookings {
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

func listRemainingTickets(){
	fmt.Printf("\n We still have %d tickets avaliable of %d tickets \n", remainingTickets, conferenceTickets);
}