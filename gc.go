package main

import "fmt"

type Book struct {
	title string
}

func main(){
	books := []*Book{}

	for i := 0; i < 10; i++ {
		book := new(Book)
		book.title = fmt.Sprintf("Book No.%d", i + 1)
		books = append(books, book)
	}

	for _, book := range books {
		fmt.Println(book.title)
	}
}