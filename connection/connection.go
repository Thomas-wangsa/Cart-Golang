package connection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"log"
	"strconv"
)

type cart struct {
	product string
	total int
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:668482@tcp(127.0.0.1:3306)/nigorice")
		if err != nil {
			os.Exit(3)
		}
	return db
}

func ORM_Get()  {
	db := Connect()
	defer db.Close()
	rows,err := db.Query("SELECT product,sum(quantity) from cart WHERE status = 1 GROUP BY product")
		if err != nil {
			log.Println(err.Error())
		}
	var result[] cart

	for rows.Next() {
		var each = cart{}
		var err = rows.Scan(&each.product,&each.total)
			if err != nil {
				log.Println(err.Error())
			}
		result = append(result, each)
	}

	for _,val := range result {
		log.Println(val.product+"("+strconv.Itoa(val.total)+")")
	}
}

func ORM_Add(product string,quantity int)  {
	db := Connect()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO cart (product,quantity) VALUES(?,?) ")
	stmt.Exec(product, quantity)
		if err != nil {
			log.Println(err.Error())
		}
}

func ORM_Delete(product string)  {
	db := Connect()
	defer db.Close()
	if _, err := db.Exec("update cart set status = ? where product = ?", 0, product); err != nil {
			log.Println(err.Error())
		}
}

func ORM_Truncate() {
	db := Connect()
	defer db.Close()
	if _, err := db.Exec("TRUNCATE TABLE cart"); err != nil {
		log.Println(err.Error())
	}
}