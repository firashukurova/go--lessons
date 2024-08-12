package main

import "fmt"

//2 Описание: Реализуйте структуру Student с полем grades (срез оценок). Реализуйте метод AddGrade, добавляющий оценку, и метод AverageGrade,
//возвращающий среднее значение оценок.
//Методы:
//AddGrade(grade int)
//AverageGrade() float64

type Student struct {
	Grades []int
}

func (s *Student) AddGrad(grade int) {
	s.Grades = append(s.Grades, grade)
}

func (s *Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

func main() {
	student := Student{}
	student.AddGrad(45)
	student.AddGrad(55)
	student.AddGrad(60)
	student.AddGrad(33)

	fmt.Printf("Средняя оценка студента: %v\n", student.AverageGrade())
}
