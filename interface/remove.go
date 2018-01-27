package Interface

import (
	Connection "Cart-Golang/connection"
)

// Interface Remove Product
func (d DepedencyInjection) RemoveProduct(product string) {
	Connection.ORM_Delete(product)
}
