package main

import (
	"fmt"
)

// Defining Grade Enum using iota
type Grade int

const (
	F = iota
	C
	B
	A
)

var majors = []string{"Computer Science", "Mathematics", "Physics", "Biology", "Chemistry", "Engineering", "Economics", "Business", "Political Science", "History"}

// Function to get a pointer to a major
// Otherwise, we will need to create a variable , assign value to it, and make student major point to it
func getMajorPointer(index int) *string {
	if index < 0 || index >= len(majors) {
		return nil // Return nil if index is out of bounds
	}
	return &majors[index]
}

// String method to convert enum to a string
func (g Grade) String() string {
	return [...]string{"F", "C", "B", "A"}[g] //[...] is an array literal that allows you to not specify the length of the array
}

type person struct {
	name string
	age  int
}

type student struct {
	person
	enrolledInCollege bool
	major             *string //This is a pointer to a string, which means it can be nil (NULLABLE)
	minor             string  //However, it is not required because you can check if the value is empty or not(for strings "", for int 0)
	grade             Grade
}

func (p person) isAdult() bool { //a method that takes a person and returns a boolean
	return p.age >= 18
}

func (p person) greet() {
	fmt.Println("Hello, my name is", p.name)
}

func (s student) greet() {
	if s.enrolledInCollege {
		fmt.Println("Hello, my name is", s.name, "and I am a student", "majoring in", *s.major, "with a grade of", s.grade)
	} else {
		fmt.Println("Hello, my name is", s.name, "and I am not a student")
	}
}

func main() {
	// Use a pointer to the major

	p := person{name: "John", age: 30}
	p.greet()
	s := student{person: person{name: "Jane", age: 20}, enrolledInCollege: true, major: getMajorPointer(2), grade: A}
	s.greet()
	sp := student{person: p, enrolledInCollege: false}
	sp.greet()
}
