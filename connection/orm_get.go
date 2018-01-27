package connection

import (
	"log"
	"strconv"
)

// ORM GET active Product
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

