/*
go mod init <application-name> => the first thing is run this command line
It creates go.mod, which describes the project we have on the IDE
You will always need to put your application file on  package
To do it, you can write "package <package-name>"
*/


package main
/*
You need say where the application begins to the compiler.
This entry function is always the func main(){}
For each file, you must have only one main() function.
*/

/*
When we download Go settings, it already gives to us some packages we can use,
like fmt, but you may have to install another, depending on your application
Each package has functions we can use. 
It's like a container with lots of functionalities
*/

import "fmt"
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
var bookings []string; //its a slice, which is almost the same as an array
const conferenceTickets = 50;
var remainingTickets uint = 50; //variable typed like uint are not allowed to recieve negative values
var userName string;
var userTickets uint;

func getUserName() {
	fmt.Printf("Please, tell me your name: ");
	fmt.Scanf("%v", &userName); 
}

func getUserTickets() {
	fmt.Printf("\nHow many tickets do you wish to buy? Answer: ");
	fmt.Scanf("%d", &userTickets); //o & serve para o compilador também ler o endereço de memória da variável userTickets
}

func validateUserInput() int{
	var userNameLenght = len(userName);
	return userNameLenght;
}

func buyTickets(userNameLenght int){
	if (userNameLenght > 2) {
		if (userTickets > remainingTickets) {
			fmt.Printf("\nSorry, we only have %d tickets left \n", remainingTickets);
		} else {
			fmt.Println("Thanks for helping us!");
			remainingTickets = remainingTickets - userTickets;
			bookings = append(bookings, userName);
		}
	} else {
		fmt.Println("Your name is not valid! Please, try again");
		getUserName();	
		var lenght = validateUserInput();
		buyTickets(lenght);
	}
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

func endSystem(){
	fmt.Println("Thanks, goodbye!");
}

func main() {	
	var conferenceName = "Go Conference";
	var wish int;

	
	//bookings := [49] string{};
	//var bookings = [50]string{"Nana", "Nicole", "Peter"};

	fmt.Printf("Welcome to %s booking application \n", conferenceName);

	//in Go, we only have the for loop
	for (wish != 4 && remainingTickets != 0) {
		fmt.Println("================> MENU <================");
		fmt.Printf("1 - Buy tickets \n2 - List Buyers \n3 - List Remaining Tickets \n4 - Exit System \n : ");
		fmt.Scan(&wish)

		switch wish {
			case 1:
				getUserName();
				getUserTickets();
				var length = validateUserInput();
				buyTickets(length);
			case 2:
				listBuyers();
			case 3:
				listRemainingTickets();
			default:
				endSystem();
		}
	}
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
