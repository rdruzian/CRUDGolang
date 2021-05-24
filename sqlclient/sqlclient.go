package sqlclient

import (
	"awesomeProject/dbConn"
	"awesomeProject/news"
	"fmt"
	_ "awesomeProject/news"
)

// CreateDatabase realiza a criação do banco de dados
func CreateDatabase() error {
	db, err := dbConn.DbConn()
	if err != nil {
		fmt.Println("Erro ao iniciar a conexão com o banco de ddados")
		return fmt.Errorf("Erro ao criar conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`create database if not exists News;`)
	if err != nil {
		fmt.Println("Erro ao criar banco de dados: %v", err)
		return fmt.Errorf("Erro ao criar banco de dados", err)
	}

	return nil
}

// CreateTableNews cria as tabelas no banco de dados
func CreateTableNews() error {
	db, err := dbConn.DbConn()
	if err != nil {
		fmt.Println("Erro ao iniciar a conexão com o banco de ddados")
		return fmt.Errorf("Erro ao criar conexão com o banco de dados", err)
	}
	defer db.Close()

	_, err = db.Exec(`USE News;`)
	if err != nil {
		fmt.Println("Erro ao setar banco de dados: %v", err)
		return fmt.Errorf("Erro ao setar banco de dados", err)
	}

	_, err = db.Exec(`create table if not exists Table_News(id integer not null auto_increment primary key, title varchar(150) not null, category varchar(20), text TEXT not null);`)
	if err != nil {
		fmt.Println("Erro ao criar tabela: %v", err)
		return fmt.Errorf("Erro ao criar tabela", err)
	}

	return err
}

func InsertDataNews() error {
	db, err := dbConn.DbConn()
	if err != nil {
		fmt.Println("Erro ao iniciar a conexão com o banco de dados")
		return fmt.Errorf("Erro ao criar conexão com o banco de dados", err)
	}
	defer db.Close()

	dataNews, err := news.GetValues()
	if err != nil {
		fmt.Println("Erro ler valores do json %v", err)
		return err
	}
	for _, data := range dataNews {
		_, err = db.Query(`INSERT INTO Table_News(title, category, text) VALUES (?, ?, ?)`, data.Title, data.Category, data.News)
		if err != nil {
			fmt.Println("Erro ao inserir dados na tabela notícias: %v", err)
			return fmt.Errorf("Erro ao inserir dados na tabela tabela", err)
		}
	}
	return err
}