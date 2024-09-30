// basic syntax for creating a map
// map[keyType]valueType
// To create a map, use the make()

package main

import "fmt"

type Vertex struct {
	Lat, Long float64 //fields that are of type float64
}

var m map[string]Vertex //makes it a global variable and init it to zero

func main() {
	// Initialize the map
	m = make(map[string]Vertex)

	// Insert an element
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	m["Google"] = Vertex{37.42202, -122.08408}

	// Retrieve an element
	googleLocation := m["Google"]
	fmt.Println("Google Location:", googleLocation)

	// Delete an element
	delete(m, "Bell Labs")

	// Test for presence of a key
	if elem, ok := m["Bell Labs"]; ok {
		fmt.Println("Bell Labs exists:", elem)
	} else {
		fmt.Println("Bell Labs does not exist")
	}

	// Demonstrate the zero value for a non-existent key
	elem, ok := m["Facebook"]
	if !ok {
		fmt.Println("Facebook does not exist:", elem) // elem will be Vertex{0, 0}
	}

	// Another example of updating an existing element
	m["Google"] = Vertex{37.42206, -122.08409} // Update Google's location
	fmt.Println("Updated Google Location:", m["Google"])
}
