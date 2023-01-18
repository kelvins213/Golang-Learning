package server

type Server struct{
 	Port string;
}

func Main() Server{

	Server := Server{Port: "8080"}

	return Server;
}