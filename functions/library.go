package main

import (
	"fmt"
)

// Book struct definition
type Book struct {
	Title  string
	Author string
	Rating int
}

// String method for Book
func (book Book) String() string {
	return fmt.Sprintf("Title: %s, Author: %s, Rating: %d", book.Title, book.Author, book.Rating)
}

// DigitalLibrary struct definition
type DigitalLibrary struct {
	Books *[]Book
}

// Method to add a book to the library
func (library *DigitalLibrary) addBook(book Book) {
	*library.Books = append(*library.Books, book)
}

// Method to display books in the library
func (library *DigitalLibrary) displayBooks() {
	for _, book := range *library.Books {
		fmt.Println(book)
	}
}

// Combined delete method for title and author
func (library *DigitalLibrary) deleteBook(title string, author string) {
	bookFound := false

	for index, book := range *library.Books {
		if book.Title == title || book.Author == author {
			library.deleteBookByIndex(index)
			fmt.Printf("Book with title '%s' or written by '%s' deleted\n", title, author)
			bookFound = true
			break
		}
	}

	if !bookFound {
		fmt.Printf("Book '%s' or by author '%s' not found\n", title, author)
	}
}

// Method to delete a book by index
func (library *DigitalLibrary) deleteBookByIndex(index int) {
	*library.Books = append((*library.Books)[:index], (*library.Books)[index+1:]...) // Corrected line
}
