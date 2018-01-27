package Interface

import (
	Connection "Cart-Golang/connection"
)

func RemoveProduct(product string) {
	Connection.ORM_Delete(product)
}
