package getters;

import (
	"fmt"
);

var UserName string;
var UserTickets uint;

func GetUserName() {
	fmt.Printf("Please, tell me your name: ");
	fmt.Scanf("%v", &UserName); 
}

func GetUserTickets() {
	fmt.Printf("\nHow many tickets do you wish to buy? Answer: ");
	fmt.Scanf("%d", &UserTickets); //o & serve para o compilador também ler o endereço de memória da variável userTickets
}