package main

import "fmt"

type Bmi float64
type Weight float64
type Grade int

//type BloodPressure float64

type BloodPressure struct {
	Upper float64
	Lower float64
}

func main() {
	var mass float64
	var height float64
	fmt.Println("Введите Ваш вес в кг")
	fmt.Scanln(&mass)
	fmt.Println("Введите Ваш рост в м")
	fmt.Scanln(&height)

	bmi := CalculateBMI(mass, height)
	result := CheckBMI(bmi)

	fmt.Printf("Ваш BMI: %.2f\n", bmi)
	fmt.Printf("Категория: %s\n", result)

	var mass1 float64
	fmt.Println("Введите вес в кг:")
	fmt.Scanln(&mass1)

	weight := Weight(mass1)
	category := CheckWeight(weight)

	fmt.Printf("Вес: %.2f кг\n", weight)
	fmt.Printf("Категория веса: %s\n", category)

	var grade Grade
	fmt.Println("Введите оценку")
	fmt.Scanln(&grade)
	res := CheckGrade(grade)
	fmt.Printf("Ваша оценка: %s\n", res)

	var bmiValue float64
	var upper, lower float64

	fmt.Println("Введите Ваш ИМТ: ")
	fmt.Scanln(&bmiValue)

	fmt.Println("Введите Ваше кровянное давление:")
	fmt.Scanf("%d %d", &upper, &lower)

	bmi = Bmi(bmiValue)
	bp := BloodPressure{Upper: upper, Lower: lower}

	status := CheckHealthStatus(bmi, bp)
	fmt.Printf("Состояние здоровья: %s\n", status)

}

//Создайте псевдоним для типа float64 под названием Weight. Напишите функцию, которая принимает Weight и
//возвращает сообщение о том, в какой категории веса находится объект: "Light", "Medium" или "Heavy".
//

func CheckWeight(weight Weight) string {
	switch {
	case weight < 50:
		return "Light"
	case weight >= 50 && weight < 100:
		return "Medium"
	default:
		return "Heavy"
	}
}

//Определите тип BMI на основе float64. Напишите функцию, которая принимает значение BMI и возвращает сообщение о том,
//в какой категории оно находится: "Underweight", "Normal weight", "Overweight", или "Obesity"

func CheckBMI(bmi Bmi) string {
	switch {
	case bmi < 18.5:
		return "Underweight"
	case bmi >= 18.5 && bmi < 25:
		return "Normal weight"
	case bmi >= 25 && bmi < 30:
		return "Overweight"
	default:
		return "Obesity"
	}
}

func CalculateBMI(mass float64, height float64) Bmi {
	return Bmi(mass / (height * height))
}

//Создайте псевдоним для типа int под названием Grade. Напишите функцию, которая принимает Grade и возвращает сообщение о том,
//является ли оценка удовлетворительной (50 и выше) или нет

func CheckGrade(grade Grade) string {
	switch {
	case grade >= 50:
		return "Satisfactory"
	default:
		return "Not satisfactory"

	}
}

//Создайте псевдонимы для типов float64 под названиями BMI и BloodPressure. Напишите функцию, которая принимает BMI и BloodPressure
//и возвращает сообщение о состоянии здоровья: "Healthy", "At risk" или "Unhealthy"

func CheckHealthStatus(bmi Bmi, bp BloodPressure) string {
	if bmi < 18.5 && bmi >= 30 || bp.Upper >= 140 && bp.Lower >= 90 {
		return "Unhealthy"
	} else if (bmi >= 25 && bmi < 30) || (bp.Upper >= 120 && bp.Upper < 140) || (bp.Lower >= 80 && bp.Lower < 90) {
		return "At risk"
	} else {
		return "Healthy"
	}

}
