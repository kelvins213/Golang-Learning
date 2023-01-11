/*
go mod init <application-name> => the first thing is run this command line
create go.mod, which describes the project we have on the IDE
you will always need to put your application file on  package
to do it, you can write "package <package-name>"
*/


package main
/*
you need say to wherethe application begins to the compiler.
this entry function is always the func main(){}
for each file, you must have only one main() function.
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

func main() {
	fmt.Print("Hello Golang bitches \n")
	fmt.Printf("Hello bitches and bros and nonbinary hoes \n");
	fmt.Println("Hello world finally");
}

/*
To the application, you can run:
go run <file name>
*/
