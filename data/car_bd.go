package people

import (
	people "booking-app/domain"
	"fmt"
)

func CreateCars(carName string, manufactory string, model string) people.Car{

	var car people.Car = people.Car{
		Name: carName,
		Manufacturer: manufactory,
		Model: model,
	};

	fmt.Println(car.Manufacturer);
	
	return car;
}