package api

import (
	db "awesomeProject/dbConn"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple project!")
}

func Show (w http.ResponseWriter, r *http.Request) {
	var (
		title, category, text string
	)
	db:= db.DbConn()
	defer db.Close()

	news, err := db.Query("select * from News")
	if err != nil {
		fmt.Println("Error to get news on database %v", err)
	}
	for news.Next() {
		err = news.Scan(&title, &category, &text)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Fprintln(w, title)
		fmt.Fprintln(w, category)
		fmt.Fprintln(w, text)
	}
}