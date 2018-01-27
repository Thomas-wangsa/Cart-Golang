package main

import (
    Interface "Cart-Golang/interface"
)

func main() {
    var interfaces Interface.Interfaces
    interfaces = Interface.DepedencyInjection{}

    interfaces.RefreshData()
    interfaces.AddProduct("Product 1",2)
    interfaces.AddProduct("Product 1",5)
    interfaces.AddProduct("Product 2",2)
    interfaces.AddProduct("Product 4",2)
    interfaces.RemoveProduct("Product 2")
    interfaces.Show_Cart()
}

