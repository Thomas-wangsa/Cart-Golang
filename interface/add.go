package Interface

import (
	Connection "Cart-Golang/connection"
)

// Interface Add Product
func (d DepedencyInjection) AddProduct(product string,quantity int) {
	Connection.ORM_Add(product,quantity)
}
