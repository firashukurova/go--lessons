package main

//Описание: Реализуйте структуру Student с полем name и age. Реализуйте структуру Classroom с срезом студентов и методы для добавления студента
//и получения средней возрастной группы.
//Методы:
//AddStudent(name string, age int)
//AverageAge() float64

type Student1 struct {
	Name string
	Age  int
}

type Classroom struct {
	Student1
}

func (s *Student1) AddStudent() {
	
}

func main() {

}
