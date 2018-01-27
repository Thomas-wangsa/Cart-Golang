package Interface

import (
	Connection "Cart-Golang/connection"
)

// Interface Show Cart
func (d DepedencyInjection) Show_Cart()  {
	Connection.ORM_Get()
}
