package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
);

type User struct{
	Id int;
	Name string;
}

func main(){
	http.HandleFunc("/users", getUsers);
	fmt.Println("api is running on: http://localhost:8080");
	log.Fatal(http.ListenAndServe(":8080", nil)); //configura a porta que o servidor irá rodar
}

func getUsers(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content=Type", "application/json"); //configura a resposta como um json
	json.NewEncoder(response).Encode(User{
		Id: 100000,
		Name: "Kelvin",
	});
}
