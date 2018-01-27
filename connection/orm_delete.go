package connection

import "log"

func ORM_Delete(product string)  {
	db := Connect()
	defer db.Close()
	if _, err := db.Exec("update cart set status = ? where product = ?", 0, product); err != nil {
		log.Println(err.Error())
	}
}