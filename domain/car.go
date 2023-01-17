package people

import (
	"fmt"
);

type Car struct{
	Name string;
	Manufacturer string;
	Model string;
}

func (c Car) ListCar(){
	fmt.Printf("Name: [%s]\n", c.Name);
	fmt.Printf("Modelo: [%s]\n", c.Model);
	fmt.Printf("Fabricante: [%s]\n", c.Manufacturer);	
}