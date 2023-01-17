package people

import (
	"fmt"
	"booking-app/domain"
);


var humans [] people.Person;

func Main(name string, age uint, car people.Car){

	/*
		someone := people.Person{
			Name: name,
			Age: age,
		};
		fmt.Println(someone.Name);
	*/
	
	person := people.Person{Name: name, Age: age, Car: car};
	person.ListPerson();

	humans = append(humans,  people.Person{Name: name, Age: age, Car: car});	

}

func ListPeople(){

	for _, human := range humans{
		fmt.Printf("[%s]:[%d] \n", human.Name , human.Age);
		
		fmt.Println("Carro");
		fmt.Printf("Nome: [%s] \n", human.Car.Name);
		fmt.Printf("Modelo: [%s] \n", human.Car.Model);
		fmt.Printf("Fabricante: [%s] \n", human.Car.Manufacturer);
	}
	
}