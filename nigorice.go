package main

import (
    Interface "Cart-Golang/interface"
)

func main() {
    Interface.RefreshData()
    Interface.AddProduct("Product 1",2)
    Interface.AddProduct("Product 1",5)
    Interface.AddProduct("Product 2",2)
    Interface.AddProduct("Product 4",2)
    Interface.RemoveProduct("Product 2")
    Interface.Show_Cart()
}

