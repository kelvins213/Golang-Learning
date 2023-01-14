/*
go mod init <application-name> => the first thing is run this command line
It creates go.mod, which describes the project we have on the IDE
You will always need to put your application file on  package
To do it, you can write "package <package-name>"
*/


package main
/*
You must say where the application begins to the compiler.
This entry function is always the func main(){}
For each file, you must have only one main() function.
*/

/*
When we download Go settings, it already gives us some packages we can use,
like fmt, but you may have to install another, depending on your application
Each package has functions we can use. 
It's like a container with lots of functionalities
*/


import (
	"fmt"
	"booking-app/geters"
	"booking-app/validation"
	"booking-app/listing"
	"booking-app/helper"
);
/*
the object fmt has the basics go's functions, like Print and Scanf 
*/

/*
we can import more than a package doing like this:
import (
	"fmt"
	"strings"
);
*/
 //its a slice, which is almost the same as an array
const ConferenceTickets = 50;
//variable typed like uint are not allowed to recieve negative values

func ShowMenu(){
	var wish int;

	
	//bookings := [49] string{};
	//var bookings = [50]string{"Nana", "Nicole", "Peter"};


	//in Go, we only have the for loop
	for (wish != 4 && RemainingTickets != 0) {
		fmt.Println("================> MENU <================");
		fmt.Printf("1 - Buy tickets \n2 - List Buyers \n3 - List Remaining Tickets \n4 - Exit System \n: ");
		fmt.Scan(&wish)

		switch wish {
			case 1:
				userInteration();
			case 2:
				listing.ListBuyers();
			case 3:
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
	geters.GetUserName();
	geters.GetUserTickets();
	validation.ValidateUserInputs();
}


func main() {	
	var conferenceName = "Go Conference";
	fmt.Printf("Welcome to %s booking application \n", conferenceName);
	ShowMenu();
	//iteration on a list
	/*
		for index, nameBooking := range bookings {

		}
	*/
	/*
	infinites loops on Golang
		for {

		}
		for true {

		}
	*/



	//fmt.Printf("conferenceTickets is %T, remaingTickets is %T, conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName);

	//fmt.Printf("How many tickets do you wish to buy? Answer: ");
	//fmt.Scanf("%d", &userTickets); //o & serve para o compilador também ler o endereço de memória da variável userTickets

	//fmt.Println(userTickets);
	//fmt.Println(&userTickets); //prints the memory addres which points to userTickets value

}

/*
To build the application, you can run:
go run <file name>
*/
