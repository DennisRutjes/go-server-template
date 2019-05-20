package main

import (
	"flag"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"server-template/impl"
	"server-template/server"
)

func main() {

	bindHost := flag.String("s", "0.0.0.0", "server internal bind address, default; 0.0.0.0")
	bindPort := flag.Int("p", 8080, "server internal bind port, default; 8080")
	serverEndpoint := flag.String("endpoint", "http://localhost:8080", "server endpoint for proxy reasons, default; http://localhost:8080")

	flag.Parse()

	service := new(impl.SayHello)

	apiServer, err := server.NewServer(service)

	if err != nil {
		log.Fatal("Server initialisation error:", err)
	}

	router := apiServer.Router()
	handler := cors.Default().Handler(router)

	// ServerEndpoint can be different from bindaddress due to routing externally
	bindAddress := fmt.Sprintf("%v:%v", *bindHost, *bindPort)

	log.Print("|")
	log.Printf("| SERVING ON: %s", *serverEndpoint)

	if err := http.ListenAndServe(bindAddress, handler); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
