package sql

import (
	"awesomeProject/dbConn"
	"fmt"
	"io/ioutil"
)

// CreateDatabase realiza a criação do banco de dados
func CreateDatabase() error {
	db, err := dbConn.DbConn()
	if err != nil {
		fmt.Println("Erro ao iniciar a conexão com o banco de ddados")
		return fmt.Errorf("Erro ao criar conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	file, err := ioutil.ReadFile("./createDatabase.sql")

	if err != nil {
		fmt.Println("Erro ao ler arquivo de criação do banco de dados %v", err)
		return fmt.Errorf("Erro ao ler arquivo sql %v", err)
	}

	criaBancoDados := string(file)

	_, err = db.Exec(criaBancoDados)
	if err != nil {
		fmt.Println("Erro ao criar banco de dados: %v", err)
		return fmt.Errorf("Erro ao criar banco de dados: %v", err)
	}

	return nil
}

func CreateTableNews() error {
	db, err := dbConn.DbConn()
	if err != nil {
		fmt.Println("Erro ao iniciar a conexão com o banco de ddados")
		return fmt.Errorf("Erro ao criar conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	file, err := ioutil.ReadFile("./createDatabase.sql")

	if err != nil {
		fmt.Println("Erro ao ler arquivo de criação da tabela %v", err)
		return fmt.Errorf("Erro ao ler arquivo sql %v", err)
	}

	criaTabelaNoticias := string(file)

	_, err = db.Exec(criaTabelaNoticias)
	if err != nil {
		fmt.Println("Erro ao criar tabela: %v", err)
		return fmt.Errorf("Erro ao criar tabela: %v", err)
	}

	return err
}