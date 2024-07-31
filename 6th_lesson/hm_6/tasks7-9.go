package main

import "fmt"

func main() {
	expenses := 0.0

	for {
		var sum float64
		fmt.Println("Введите сумму")
		fmt.Scan(&sum)

		if sum == 0 {
			break
		}

		AddExpenses(&expenses, sum)
		fmt.Println()

	}
}

//Учет ежедневных расходов с лимитом. Начальный лимит расходов в день равен 5000 рублей. Реализуйте функции для добавления расходов в течение дня.
//Выводите текущую сумму расходов после каждого добавления. Если сумма расходов превышает лимит, выведите сообщение "Превышен дневной лимит".

func AddExpenses(expenses *float64, sum float64) {
	*expenses += sum
	fmt.Printf("Текущая сумма расходов %.2f\n", *expenses)

	if *expenses > 5000 {
		fmt.Println("Превышена сумма расходов")
	}
}

//Начальная сумма депозита равна 500000 рублей. Реализуйте функции для начисления процентов каждый год по ставке 6%. Выводите баланс после каждого начисления.
//Если баланс становится больше 700000 рублей, выведите сообщение "Достигнут лимит начислений"
//func InterestCalculation(deposit *float64) {
//	percent := *deposit * 0.05 //5% от текущей суммы
//	*deposit += percent
//	fmt.Printf("Новый баланс: %.2f рублей\n", *deposit)
//
//	if *deposit > 300000 {
//		fmt.Println("Достигнут лимит начислений")
//	}
//
//}

func InterestCalculation2(deposit3 *float64) {

}
