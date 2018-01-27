package connection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"log"
)

type cart struct {
	product string
	total int
}

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:668482@tcp(127.0.0.1:3306)/nigorice")
		if err != nil {
			os.Exit(3)
		}
	return db
}

func ORM_GET() error {
	db := connect()
	defer db.Close()
	rows,err := db.Query("SELECT product,sum(quantity) from cart WHERE status = 1 GROUP BY product")
		if err != nil {
			return err
		}
	var result[] cart

	for rows.Next() {
		var each = cart{}
		var err = rows.Scan(&each.product,&each.total)
			if err != nil {
				return err
			}
		result = append(result, each)
	}

	for key,val := range result {
		log.Println(key)
		log.Println(val)
	}

	return nil


}

func ORM_Add(product string,quantity int) error {
	db := connect()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO cart (product,quantity) VALUES(?,?) ")
	stmt.Exec(product, quantity)
	if err != nil {
		return err
	}
	return nil
}