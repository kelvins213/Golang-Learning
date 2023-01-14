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
	"booking-app/menu"
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



func main() {	
	var conferenceName = "Go Conference";
	fmt.Printf("Welcome to %s booking application \n", conferenceName);
	menu.ShowMenu();
	



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
