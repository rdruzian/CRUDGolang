package main

import (
	"awesomeProject/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// The routes initialization must be before initialize the web server
	routes.Init()

	fmt.Println("Server is running on 9000 port!")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}