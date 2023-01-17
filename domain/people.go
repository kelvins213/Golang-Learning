package people

import (
	"fmt"
);

type Person struct{
	Name string;
	Age uint;
	Car Car;
}

func (p Person) ListPeople(){
	fmt.Printf("[%s]:[%d] \n", p.Name , p.Age);	
	fmt.Println("Carro");
	fmt.Printf("Nome: [%s] \n", p.Car.Name);
	fmt.Printf("Modelo: [%s] \n", p.Car.Model);
	fmt.Printf("Fabricante: [%s] \n", p.Car.Manufacturer);
}