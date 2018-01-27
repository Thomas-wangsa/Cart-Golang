package Interface

import (
	Connection "Cart-Golang/connection"
)

func RefreshData() {
	Connection.ORM_Truncate()
}

func Show_Cart()  {
	Connection.ORM_Get()
}

func AddProduct(product string,quantity int) {
	Connection.ORM_Add(product,quantity)
}

func RemoveProduct(product string) {
	Connection.ORM_Delete(product)
}