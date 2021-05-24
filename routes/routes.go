package routes

import (
	"awesomeProject/api"
	"net/http"
)

func Init () {
	http.HandleFunc("/", api.Home)
	//http.HandleFunc("/criabanco", api.CriaBanco)
}