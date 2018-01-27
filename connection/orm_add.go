package connection

import "log"

// ORM Add Product
func ORM_Add(product string,quantity int)  {
	db := Connect()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO cart (product,quantity) VALUES(?,?) ")
	stmt.Exec(product, quantity)
	if err != nil {
		log.Println(err.Error())
	}
}
