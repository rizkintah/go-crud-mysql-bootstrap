package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := ""
	dbName := "golang_crud"

	// open connection
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	return db, err
}
