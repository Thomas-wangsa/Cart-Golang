package Interface

import (
	Connection "Cart-Golang/connection"
)

func (d DepedencyInjection) Show_Cart()  {
	Connection.ORM_Get()
}
