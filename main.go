//Na Golang, não existem Classes, no entanto,
//existem estruturas similares, como as structs.
//

package main

import (
	"fmt"
);

var (
	people []person
	name string;
	age uint;
	wish = false;
);


//person é como um construtor, que criara um objeto
//do tipo person com os atributos de sua estrutura
type person struct{
	name string;
	age uint;
}

func showMenu(){
	for {
		fmt.Printf("Deseja cadastrar informações? [True | False]: ");
		fmt.Scan(&wish);
		if wish {
			fmt.Printf("\n Digite o nome da pessoa: ");
			fmt.Scan(&name);
			fmt.Printf("\n Digite a idade da pessoa: ");
			fmt.Scan(&age);
			people = append(people, person{name: name, age: age});		
		} else {
			break;
		}
	}
}

func listPeople(){
	for _, person := range people{
		fmt.Printf("[%s]:[%d] \n", person.name, person.age);
	}
}

func main() {
	
	showMenu();	
	
	fmt.Printf("\n Deseja listar as informações? [True | False]: \n");
	fmt.Scan(&wish);
	
	if wish {
		listPeople();	
	}
}
