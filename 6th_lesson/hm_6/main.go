package main

import "fmt"

//Начальный баланс накопительного счета равен 100000 рублей. Реализуйте функции для пополнения счета каждый месяц на фиксированную сумму.
//Выводите баланс после каждого пополнения.Если баланс становится больше 500000 рублей, выведите сообщение "Достигнут лимит накоплений"

func main() {
	MonthlyReplenishment()
	LoanBalance()

	balance := 500000.0
	for {
		var sum float64
		fmt.Print("Введите сумму перевода (0 для выхода): ")
		fmt.Scan(&sum)
		if sum == 0 {
			break
		}
		BankTransfer(&balance, sum)
	}

}

func Refill(balance *float64, sum float64) {
	*balance += sum
	fmt.Printf("\"Пополнение на %.2f руб. Новый баланс: %.2f руб.\\n\"", sum, *balance)

	if *balance > 500000 {
		fmt.Println("Savings limit reached")
	}
}
func MonthlyReplenishment() {
	balance := 100000.0
	monthlyReplenishment := 50000.0
	for month := 1; month <= 12; month++ {
		fmt.Printf("Месяц %d:\n", month)
		Refill(&balance, monthlyReplenishment)
		fmt.Println()

		if balance > 500000 {
			break
		}
	}
}

//Начальная сумма кредита равна 3000000 рублей.Реализуйте функции для ежемесячного расчета выплат по кредиту с фиксированной процентной ставкой 10%.
//Выводите остаток по кредиту после каждой выплаты.Если остаток по кредиту становится меньше 500000 рублей, выведите сообщение "Почти погашен кредит"

func CreditCalculator(balance *float64, payment float64) {
	interestRate := *balance * 0.1 / 12 // 10% годовых
	*balance = *balance + interestRate - payment

	if *balance < 0 {
		*balance = 0
	}
}

func LoanBalance() {
	initialLoanAmount := 3000000.0
	monthlyPayment := 100000.0

	for month := 1; initialLoanAmount > 0; month++ {
		fmt.Printf("Month %d:\n", month)
		CreditCalculator(&initialLoanAmount, monthlyPayment)
		fmt.Printf("Credit: %.2f руб.\n\n", initialLoanAmount)

		if initialLoanAmount < 500000 {
			fmt.Println("Loans limit reached")
		}
	}
}

//Учет операций по банковским переводам с лимитом суммы. Начальный баланс счета равен 500000 рублей. Реализуйте функции для выполнения банковских переводов.
//Если сумма перевода превышает 100000 рублей, выведите сообщение "Сумма перевода превышает лимит". Выводите остаток на счете после каждого перевода

func BankTransfer(balance *float64, sum float64) {
	if sum > 100000 {
		fmt.Println("Сумма перевода превышает лимит")
		return
	}
	*balance -= sum // Вычитаем сумму перевода
	fmt.Printf("Перевод выполнен. Остаток на счете: %.2f рублей\n", *balance)
}
