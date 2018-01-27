package Interface

import (
	Connection "github.com/Thomas-wangsa/Cart-Golang/connection"
)

// Init Interfaces
type Interfaces interface {
	RefreshData()
	AddProduct(product string,quantity int)
	RemoveProduct(product string)
	Show_Cart()
}

type DepedencyInjection struct {
}

// Interface Truncate Data
func (d DepedencyInjection) RefreshData() {
	Connection.ORM_Truncate()
}
