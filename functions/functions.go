//defer

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	functionAddder := adder() //used adder to get fucntion, and stored in functionAddder
	fmt.Println(functionAddder(1))
	display("Bye Bye")
	fmt.Println(add(1, 2)) //basic functions
	applyFunctionToContainer(appendRandomElement)

	books := []Book{
		{"The Great Gatsby", "F. Scott Fitzgerald", 5},
		{"1984", "George Orwell", 4},
		{"To Kill a Mockingbird", "Harper Lee", 5},
	}

	library := DigitalLibrary{Books: &books}

	// Display the books
	fmt.Println("Books in the library:")
	library.displayBooks()

	// Delete a book by title
	library.deleteBook("1984", "")

	// Display books after deletion
	fmt.Println("Books in the library after deletion:")
	library.displayBooks()

	// Delete a book by author
	library.deleteBook("", "Harper Lee")

	// Display books after deletion
	fmt.Println("Books in the library after deleting by author:")
	library.displayBooks()

}

func add(x int, y int) int { //(params) return type
	return x + y
}

func display(data string) { //no return type
	fmt.Println(data)
}

// higher order function that return function
func adder() func(int) int { //return type is function that takes int and returns int
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// closure
// function that captures the variables from the outside of the function
// in this case, function_applied_to_containter captures container
// and modifies it (to better illustrate closure container was shared as pointer)
func applyFunctionToContainer(function_applied_to_containter func(*[]string)) {
	container := []string{"item1", "item2", "item3"}
	function_applied_to_containter(&container)
	fmt.Println("Container after Function Applied", container)
}

func appendRandomElement(container *[]string) {
	random_string := []string{"itemX", "itemY", "itemZ"}
	*container = append(*container, random_string[rand.Intn(len(random_string))])
	fmt.Println("Container after adding random element", *container)
}
