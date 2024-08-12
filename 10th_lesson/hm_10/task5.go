package main

import "fmt"

//5 Описание: Реализуйте структуру Transaction с полями amount и description. Реализуйте метод Process, который выводит информацию о транзакции.
//Методы:
//Process()
//AddTransaction(amount float64, description string)
//Реализуйте структуру Transaction с полями amount и description. Реализуйте метод Process, который выводит информацию о транзакции.
//Реализуйте структуру Account с методом AddTransaction.
//Методы:
// Process()
// AddTransaction(amount float64, description string)

type Transaction struct {
	Amount      float64
	Description string
}
type Account struct {
	Transactions []Transaction
}

func (t *Transaction) Process() {
	fmt.Println(t.Amount, t.Description)
}

func (a *Account) AddTransaction(amount float64, description string) {
	newTransaction := Transaction{amount, description}
	a.Transactions = append(a.Transactions, newTransaction)
}

func main() {
	account := Account{}

	account.AddTransaction(100.50, "Зарплата")
	account.AddTransaction(-50.25, "Покупка продуктов")
	account.AddTransaction(30.00, "Возврат долга")

	fmt.Println("Список транзакций:")
	for _, transaction := range account.Transactions {
		transaction.Process()
	}

}
