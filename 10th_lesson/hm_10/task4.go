package main

import "fmt"

//4 Описание: Реализуйте структуру Item с полями name и price. Реализуйте структуру Inventory с срезом товаров и методы для добавления товара и получения общего
//значения товаров в инвентаре.
//Методы:
//AddItem(item Item)
//TotalValue() float64

type Item struct {
	Name  string
	Price float64
}

type Inventory struct {
	Products []Item
}

func (i *Inventory) AddItem(name string, price float64) {
	i.Products = append(i.Products, Item{name, price})
}

func (i *Inventory) TotalValue() float64 {
	totalPrice := 0.0
	for _, product := range i.Products {
		totalPrice += product.Price
	}
	return totalPrice
}

func main() {

	inventory := Inventory{}

	inventory.AddItem("Книга", 15.99)
	inventory.AddItem("Ручка", 1.99)
	inventory.AddItem("Ноутбук", 999.99)

	fmt.Printf("Общая стоимость инвентаря: %.2f сомони\n", inventory.TotalValue())

}
