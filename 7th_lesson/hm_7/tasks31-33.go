package main

import "fmt"

type Node struct {
	num  int
	next *Node
}

func main() {

	MakeNodes()
	MakeNodes3()

	head := CreateLinkedList()
	TraverseLinkedList(head)

}

//Создайте структуру Node, которая содержит целое значение и указатель на следующий узел. Напишите функцию,
//которая создает два связанных узла и выводит их значения.

func MakeNodes() {
	node1 := &Node{num: 9}
	node2 := &Node{num: 5}
	node1.next = node2

	fmt.Println("Значение первого узла:", node1.num)
	fmt.Println("Значение второго узла:", node2.num)

	fmt.Println("Значение следующего узла после первого:", node1.next.num)
}

//Создайте структуру Node, которая содержит целое значение и указатель на следующий узел. Напишите функцию,
//которая создает три связанных узла и выводит их значения

func MakeNodes3() {
	node1 := &Node{num: 1}
	node2 := &Node{num: 2}
	node3 := &Node{num: 3}

	node1.next = node2
	node2.next = node3

	fmt.Println("Значение первого узла:", node1.num)
	fmt.Println("Значение второго узла:", node2.num)
	fmt.Println("Значение третьего узла:", node3.num)
}

//Создайте структуру Node, которая содержит целое значение и указатель на следующий узел. Напишите функцию,
//которая принимает указатель на первый узел и обходит связный список, выводя значения всех узлов

func CreateLinkedList() *Node {
	node1 := &Node{num: 1}
	node2 := &Node{num: 2}
	node3 := &Node{num: 3}

	node1.next = node2
	node2.next = node3

	return node1 // Возвращаем указатель на первый узел
}

func TraverseLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Println(current.num)
		current = current.next
	}
}
