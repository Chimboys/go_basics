//defer

package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// (params) return type
func add(x int, y int) int {
	return x + y
}

// basic function with no return type
func display(data string) {
	fmt.Println(data)
}

// multiple return types + error handling
func devide(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, errors.New("cannot devide by zero")
	}
	return x / y, x % y, nil //meaning no error occured
}

// named return values
func split(sum int) (x, y int) { //x and y are named return values, no need to return them explicitly
	x = sum * 4 / 9
	y = sum - x
	return
}

// higher order functions (functions that return function)
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
	defer fmt.Println("Container after adding random element", *container)
	random_string := []string{"itemX", "itemY", "itemZ"}
	*container = append(*container, random_string[rand.Intn(len(random_string))])

}

// deferExample demonstrates the use of defer
func deferExample() {
	defer fmt.Println("This message is printed after deferExample finishes executing.")
	fmt.Println("Executing deferExample...")
	// Simulate some work
	for i := 0; i < 3; i++ {
		fmt.Printf("Working... %d\n", i+1)
	}
}

func main() {
	//pt1 basic functions
	functionAddder := adder() //used adder to get fucntion, and stored in functionAddder
	fmt.Println(functionAddder(1))
	display("Bye Bye")
	fmt.Println(add(1, 2)) //basic functions
	applyFunctionToContainer(appendRandomElement)
	deferExample()
	x, y, err := devide(12, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x, y)
	}
	fmt.Println(split(17))

	//pt2 functions and structs
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
