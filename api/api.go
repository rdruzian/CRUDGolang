package api

import (
	"awesomeProject/sqlclient"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Method not supported")
		return
	}
	criaBanco()
	fmt.Fprintf(w, "Simple project!")
}

func criaBanco() {
	err := sqlclient.CreateDatabase()
	if err != nil {
		panic(err)
	}
	fmt.Println("Banco criado com sucesso!")

	err = sqlclient.CreateTableNews()
	if err != nil {
		panic(err)
	}
	fmt.Println("Tabela criada com sucesso!")

	err = sqlclient.InsertDataNews()
	if err != nil {
		panic(err)
	}
	fmt.Println("Dados inseridos com sucesso!")
}