package connection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type cart struct {
	product string
	total int
}

func Connect() *sql.DB {
	username := "root"
	password := "secret"
	host	 := "127.0.0.1:3306"
	database := "nigorice"
	credentials := username+":"+password+"@tcp("+host+")/"+database
	db, err := sql.Open("mysql", credentials)
		if err != nil {
			os.Exit(3)
		}
	return db
}

