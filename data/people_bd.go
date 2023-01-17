package people

import (
	"fmt"
	"booking-app/domain"
);


var humans [] people.Person;

func Main(name string, age uint){

	humans = append(humans,  people.Person{Name: name, Age: age});	

}

func ListPeople(){

	for _, human := range humans{
		fmt.Printf("[%s]:[%d] \n", human.Name , human.Age);
	}
	
}