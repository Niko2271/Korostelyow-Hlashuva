package main

import (
    "fmt"
    "errors"
)

type Product struct{
    ID int
    Price float64
    Name string
    Quantity int
}

type Inventory struct{
products map[int]Product
}

func (aboba *Inventory) AddProduct(product Product){
    if aboba.products == nil {
    aboba.products = make(map[int]Product)
    }
    aboba.products[product.ID] = product
}

func (aboba *Inventory) WriteOff(productID int, quantity int) error {
    product, exists := aboba.products[productID]
    if !exists {
        return errors.New("товар не найден")
    }     
    if product.Quantity < quantity {
    return errors.New("недостаточно товара")
    }
    product.Quantity -= quantity
    aboba.products[productID] = product
 
    return nil
}

func main(){
    fasf := &Inventory{}

   fasf.AddProduct(Product{ID: 1, Name: "Симафон", Price: 30000, Quantity: 10})
    fasf.AddProduct(Product{ID: 2, Name: "Компуктер", Price: 50000, Quantity: 5})

    errorchik := fasf.WriteOff(1, 3)
    if errorchik != nil {
    fmt.Println("Ошибка списания:", errorchik)
    } else {
    fmt.Println("Списано 1 телефона")
    }

    errorchik = fasf.WriteOff(2, 10)
    if errorchik != nil {
    fmt.Println("Ошибка списания:", errorchik)
    }

    errorchik = fasf.WriteOff(5, 1)
    if errorchik != nil {     
    fmt.Println("Ошибка списания:", errorchik) 
    }
}