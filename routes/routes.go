package routes

import (
	"awesomeProject/api"
	"net/http"
)

func Init () {
	http.HandleFunc("/", api.Home)
	http.HandleFunc("/show", api.Show)
}