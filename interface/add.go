package Interface

import (
	Connection "Cart-Golang/connection"
	"log"
	"os"
)


func Show_Cart()  {
	err := Connection.ORM_GET()
	if err != nil {
		log.Println(err.Error())
		os.Exit(3)
	}
}

func AddProduct(product string,quantity int) {
	err := Connection.ORM_Add(product,quantity)
		if err != nil {
			log.Println(err.Error())
			os.Exit(3)
		}
}