package Interface

import (
	Connection "Cart-Golang/connection"
)

type Interfaces interface {
	RefreshData()
	AddProduct(product string,quantity int)
	RemoveProduct(product string)
	Show_Cart()
}

type DepedencyInjection struct {
}

func (d DepedencyInjection) RefreshData() {
	Connection.ORM_Truncate()
}
