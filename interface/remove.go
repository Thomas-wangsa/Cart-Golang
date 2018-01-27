package Interface

import (
	Connection "Cart-Golang/connection"
)

func (d DepedencyInjection) RemoveProduct(product string) {
	Connection.ORM_Delete(product)
}
