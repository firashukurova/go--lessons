package main

import "fmt"

//1 Описание: Реализуйте структуру Book с полями title и author, а также метод GetDetails, который возвращает строку с деталями книги.
//Реализуйте структуру Library с массивом книг и метод AddBook, добавляющий книгу в библиотеку.
//Методы:
//GetDetails() string для структуры Book
//AddBook(book Book) для структуры Library

type Book struct {
	Title  string
	Author string
}
type Library struct {
	Books []Book
}

func (b Book) GetDetails() string {
	return fmt.Sprintf("Title: %s, Author: %s", b.Title, b.Author)
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func main() {
	book1 := Book{Title: "Master & Margarita", Author: "M.Bulgakov"}
	book2 := Book{Title: "Soul Music", Author: "Terry Pratchett"}

	fmt.Println(book1.GetDetails())

	library := Library{}
	library.AddBook(book1)
	library.AddBook(book2)

	fmt.Printf("Library has %d books\n", len(library.Books))
}
