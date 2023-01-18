package main

import (
	"GoAPI/server"
	"GoAPI/routes"
	"fmt"
	"log"
	"net/http"
);


func main(){
	routes.Main();

	server := server.Main();

	fmt.Printf("api is running on: http://localhost:%s \n", server.Port);
	log.Fatal(http.ListenAndServe(":" + server.Port, nil)); //configura a porta que o servidor irá rodar
}


