//Na Golang, não existem Classes, no entanto,
//existem estruturas similares, como as structs.
//

package main

import (
	"fmt"
	"booking-app/data"
);

var (
	name string;
	age uint;
	carName string;
	manufacturer string;
	model string;
	wish = false;
);


//person é como um construtor, que criara um objeto
//do tipo person com os atributos de sua estrutura

func showMenu(){
	for {
		fmt.Printf("Deseja cadastrar informações? [True | False]: ");
		fmt.Scan(&wish);
		if wish {

			fmt.Printf("\n Digite o nome da pessoa: ");
			fmt.Scan(&name);
			fmt.Printf("\n Digite a idade da pessoa: ");
			fmt.Scan(&age);

			fmt.Printf("\n Digite o nome do carro da pessoa: ");
			fmt.Scan(&carName);
			fmt.Printf("\n Digite o modelo do carro da pessoa: ");
			fmt.Scan(&model);
			fmt.Printf("\n Digite o fabricante do carro da pessoa: ");
			fmt.Scan(&manufacturer);

			car := people.CreateCars(carName, manufacturer, model);
			fmt.Println(car.Manufacturer);

			people.Main(name, age, car);

			} else {
			break;
		}
	}
}

func main() {
	
	showMenu();	
	
	fmt.Printf("\n Deseja listar as informações? [True | False]: \n");
	fmt.Scan(&wish);
	
	if wish {
		people.ListPeople();	
	}
}
