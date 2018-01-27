package Interface

import (
	Connection "Cart-Golang/connection"
)

func AddProduct(product string,quantity int) {
	Connection.ORM_Add(product,quantity)
}
