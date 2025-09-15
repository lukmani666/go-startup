package basics

import (
	"errors"
	"fmt"
)

// 1. variables constants
func VariablesAndConstants() {
	var name string = "Lukman"
	var age int = 20
	country := "Nigeria" // short-hand variable

	const pi = 3.14159

	fmt.Println("Variables and Constants:")
	fmt.Println("Name:", name, "Age:", age, "Country:", country, "Pi:", pi)
	fmt.Println()
}

// 2. Data Types
func DataTypes()  {
	var s string = "Hello Go"
	var i int = 42
	var f float64 = 3.14
	var b bool = true

	fmt.Println("Data Types:")
	fmt.Println("String:", s)
	fmt.Println("Interger", i)
	fmt.Println("Float", f)
	fmt.Println("Boolean", b)
	fmt.Println()
}

// 3. Operators
func Operators() {
	a, b := 10, 2
	fmt.Println("Operators:")
	fmt.Println("Arithemetic:", a+b, a-b, a*b, a/b, a%b)
	fmt.Println("Comparison:", a == b, a != b, a > b, a < b, a >= b, a <= b)
	x, y := true, false
	fmt.Println("Logical:", x && y, x || y, !x)
	fmt.Println()
}

// 4. Input and Output
func InputOutput() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("You are", age, "years old.")
	fmt.Println()
}

// 5. Control flow (if/else, switch)
func ControlFlow()  {
	score := 75
	fmt.Print("Control Flow with if/else:")
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 70 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or below")
	}

	day := 3
	fmt.Println("\nControl Flow with switch:")
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}
	fmt.Println()
}

// 6. Loops
func Loops() {
	fmt.Println("Loops:")

	fmt.Println("For loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i, " ")
	}
	fmt.Println()

	fmt.Println("While-style loop:")
	x := 1
	for x <= 3 {
		fmt.Print(x, " ")
		x ++
	}
	fmt.Println()

	fmt.Println("Range loop:")
	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Println("Index:", index, "Value", value)
	}
	fmt.Println()
}

func Add(a int, b int) int  {
	return a + b
}

//function that return multiple values
func Divide(a, b int) (int, string) {
	if b == 0 {
		return 0, "Can't divide by zero"
	} else {
		return a / b, "Success"
	}
}

// function return (int, error)
func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func ArrayExample() {
	// Define an array of 5 integers
	var numbers [5] int

	//Assign values
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	// declare and initialize at once
	colors := [3]string{"Red", "Green", "Blue"}

	fmt.Println("Numbers", numbers)
	fmt.Println("Colors", colors)
	fmt.Println("Length of colors array:", len(colors))
}

func SlicesExample()  {
	// declare a slice
	var fruits []string

	//append values (dynamic resizing)
	fruits = append(fruits, "Apple", "Banana", "Mango")

	//intialize directly
	primes := []int{2, 3, 5, 7, 11}

	//slice from array
	letters := [5]string{"a", "b", "c", "d", "e"}
	subset := letters[1:4] // from index 1 to 3

	fmt.Println("Fruits:", fruits)
	fmt.Println("Primes:", primes)
	fmt.Println("Subset of letters:", subset)
}

func MapsExample()  {
	// create a map with string keys and int values
	ages := map[string]int {
		"Lukman": 20,
		"Aisha": 25,
		"Michael": 30,
	}

	// Add a new entry
	ages["Fatima"] = 22

	//access value
	fmt.Println("Lukman's age:", ages["Lukman"])

	//check if key exists
	val, ok := ages["John"]
	if ok {
		fmt.Println("John's age:", val)
	} else {
		fmt.Println("John not found.")
	}

	//delete a key
	delete(ages, "Aisha")

	fmt.Println("All Ages:", ages)
}

func LoopExample()  {
	numbers := []int{10, 20, 30, 40}

	fmt.Println("Looping through slice:")
	for i, num := range numbers {
		fmt.Println("Index:", i, "Value", num)
	}

	ages := map[string]int{"Ali": 21, "Bola": 19, "Chika": 23}

	fmt.Println("\nLooping through map:")
	for name, age := range ages {
		fmt.Println(name, "is", age, "years old")
	}
}
