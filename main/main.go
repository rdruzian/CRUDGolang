package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is running on 9000 port!")
	http.ListenAndServe(":9000", nil)

	//routes.Init()
}