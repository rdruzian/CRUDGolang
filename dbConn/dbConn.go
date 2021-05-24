package dbConn

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DbConn() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/")
	if err != nil {
		panic(err.Error())
		return db, err
	}
	return db, nil
}