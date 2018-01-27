package Interface

import (
	Connection "github.com/Thomas-wangsa/Cart-Golang/connection"
)

// Interface Show Cart
func (d DepedencyInjection) Show_Cart()  {
	Connection.ORM_Get()
}
