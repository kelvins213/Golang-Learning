package routes

import (
	"GoAPI/model"
	"encoding/json"
	"net/http"
);

func Main(){
	setRoutes();
}

func setRoutes(){
	http.HandleFunc("/", sayHello);
	http.HandleFunc("/users", getUsers);
}

func sayHello(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content=Type", "application/json"); //configura a resposta como um json
	json.NewEncoder(response).Encode(model.User{
		ID: 1,
		Name: "Welcome to my new page",
	});
}

func getUsers(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content=Type", "application/json"); //configura a resposta como um json
	json.NewEncoder(response).Encode(
		[]model.User{
			{ID: 50000,
				Name: "Júlio",},
			{ID: 100000,
				Name: "Kelvin",},
			{ID: 3000,
				Name: "Daniel",},
		},
	);
}
