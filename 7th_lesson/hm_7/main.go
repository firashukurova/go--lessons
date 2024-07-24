package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	InStock  bool
	Quantity int
}

type Owner struct {
	FirstName string
	LastName  string
}
type BankAccount struct {
	AccountNumber string
	Balance       float64
	IsActive      bool
	Owner         Owner
}

func main() {

	var defaultProduct Product
	fmt.Println("defaultProduct:")
	fmt.Printf("%+v\n", defaultProduct)
	fmt.Printf("Валидность %v\n\n", IsValidProduct(defaultProduct))

	validProduct := Product{
		Name:     "Milk",
		Price:    100,
		InStock:  true,
		Quantity: 9,
	}
	fmt.Println("validProduct:")
	fmt.Printf("%+v\n", validProduct)
	fmt.Printf("Валидность: %v\n\n", IsValidProduct(validProduct))

	partialProduct := Product{
		Name:  "Bread",
		Price: 50,
	}

	fmt.Println("Частично заполненный продукт:")
	fmt.Printf("%+v\n", partialProduct)
	fmt.Printf("Валидность: %v\n", IsValidProduct(partialProduct))

	account := BankAccount{
		AccountNumber: "123",
		Balance:       100,
		IsActive:      true,
		Owner: Owner{
			FirstName: "James",
			LastName:  "Bond",
		},
	}
	fmt.Println(account)
}

func IsValidProduct(product Product) bool {
	return product.Name != "" && product.Price > 0 && product.Quantity > 0
}

//Напишите функцию `InitializeAccount`, которая принимает номер счета и имя владельца, и возвращает инициализированный `BankAccount` с следующими правилами:
//Баланс по умолчанию должен быть 0 Счет должен быть активным по умолчанию Если номер счета не предоставлен, используйте "0000"

func InitializeAccount(accountNumber string, firstName, lastName string) BankAccount {
	if accountNumber == "" {
		accountNumber = "0000"
	}
	return BankAccount{
		AccountNumber: accountNumber,
		Balance:       0,
		IsActive:      true,
		Owner: Owner{
			FirstName: firstName,
			LastName:  lastName,
		},
	}
}
