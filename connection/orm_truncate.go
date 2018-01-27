package connection

import "log"

// ORM Truncate
func ORM_Truncate() {
	db := Connect()
	defer db.Close()
	if _, err := db.Exec("TRUNCATE TABLE cart"); err != nil {
		log.Println(err.Error())
	}
}
