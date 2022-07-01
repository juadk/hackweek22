package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/juadk/hackweek22/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)

	fmt.Printf("Starting application on port %s\n", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal(err)
	}

}
