package Interface

import (
	Connection "Cart-Golang/connection"
)

func RefreshData() {
	Connection.ORM_Truncate()
}
