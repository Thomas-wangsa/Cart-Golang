package main

import (
    Interface "Cart-Golang/interface"
)

func main() {
    Interface.RefreshData()
    Interface.AddProduct("Thomas",2)
    Interface.AddProduct("Thomas",2)
    Interface.AddProduct("Thomas",2)
    Interface.AddProduct("Thomas2",2)
    Interface.RemoveProduct("Thomas2")
    Interface.Show_Cart()
}

