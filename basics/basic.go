package basics

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
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

//defining a struct
type Person struct {
	Name string
	Age int
	City string
}

func StructExample() {
	//create a struct using field names
	p1 := Person{Name: "Lukman", Age: 20, City: "Lagos"}

	//create a struct without field names (order matters)
	p2 := Person{"Aisha", 25, "Abuja"}

	//zero value struct
	var p3 Person

	fmt.Println("Person 1:", p1)
	fmt.Println("Person 2:", p2)
	fmt.Println("Person 3 (zero value):", p3)
}

// Methods on Structs
//Go doesn’t have classes, but you can define methods on structs.
//method with value receiver
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name, "and I live in ", p.City)
}

//method with value receiver that returns something
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func MethodExample() {
	p := Person{Name: "Micheal", Age: 30, City: "Ibadan"}
	p.Greet()
	fmt.Println("Is Adult?", p.IsAdult())
}

// method with pointer receiver (can modify struct)
func (p *Person) UpdateCity(newCity string) {
	p.City = newCity
}

func PointerExample() {
	p := Person{Name: "Fatima", Age: 22, City: "Kano"}
	fmt.Println("Before update:", p)

	//update city using pointer receiver
	p.UpdateCity("Abuja")

	fmt.Println("After update:", p)
}

//define an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

//define a struct
type Rectangle struct {
	Width, Height float64
}

//rectangle implements Shape
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func InterfaceExample() {
	var s Shape
	s = Rectangle{Width: 10, Height: 5}

	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

//polymorphism in GO
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return  2 * 3.14 * c.Radius
}

func PolymorphismExample() {
	shapes := []Shape{
		Rectangle{Width: 4, Height: 6},
		Circle{Radius: 3},
	}

	for _, shapes := range shapes {
		fmt.Println("Area:", shapes.Area(), "Perimeter:", shapes.Perimeter())
	}
}

//importing standard library packages
func StdLibExample() {
	fmt.Println("Square root of 16:", math.Sqrt(16))
	fmt.Println("Uppercase:", strings.ToUpper("lukman"))
}


//working with goroutine
func PrintMessage(msg string) {
	for i := 0; i < 3; i ++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func GoroutineExample() {
	//run this in a goroutine
	go PrintMessage("From Goroutine")
	
	//run in main thread
	PrintMessage("From Main Funtion")

	// give goroutine time to finish
	time.Sleep(2 * time.Second)
}

func ChannelExample() {
	ch := make(chan string)

	// sender goroutine
	go func ()  {
		ch <- "Hello from goroutine!" // send data
	} ()

	//receiver
	message := <- ch // receive data
	fmt.Println("Received:", message)
}

func SelectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	//goroutine 1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func ()  {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()

	//wait for whichever comes first
	select {
	case msg1 := <- ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <- ch2:
		fmt.Println("Received:", msg2)
	case <- time.After(3 * time.Second):
		fmt.Println("Timeout: no message received")
	}
}

// real-world example: Worker pool
func worker(id int, jobs <- chan int, results chan <- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "processing job", j)
		time.Sleep(time.Second) // simulate work
		results <- j * 2
	}
}

func WorkPoolExample() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	//collect results
	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results)
	}
}

// worker simulates a task done by a goroutine
func secondWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done() //mark this worker as done once finished
	fmt.Printf("Worker %d starting\n", id)

	//simulate some work (like processing data)
	for i := 1; i <= 3; i++ {
		fmt.Printf("Worker %d working on step %d\n", id, i)
	}

	fmt.Printf("Worker %d finished\n", id)
}

func SyncWorker() {
	var wg sync.WaitGroup //create a WaitGroup

	// launch 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1) // add one worker to WaitGroup
		go secondWorker(i, &wg) // run worker as goroutine
	}

	//wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers completed ✅")
}

func workerChannel(id int, wg *sync.WaitGroup, results chan <- string) {
	defer wg.Done() // signal completion when function ends

	//simulate some work
	message := fmt.Sprintf("Worker %d finished its job", id)

	// send result into channel
	results <- message
}

func SyncWorkerChan() {
	var wg sync.WaitGroup
	results := make(chan string, 5) // buffered channel (size 5)

	//start 3 workers
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go workerChannel(i, &wg, results)
	}

	//launch a separate goroutine to close the channel
	// once all workers are done
	go func() {
		wg.Wait() // wait for all workers
		close(results) // close the channel
	}()

	// receive results from the channel
	for results := range results {
		fmt.Println(results)
	}

	fmt.Println("All workers completed")
}

func WriteFile() {
	//create or overwrite file
	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close() // ensure file closes when function exists

	//write text to file
	_, err = file.WriteString("Hello, Go File Handling! 🚀\n")
	if err != nil {
		panic(err)
	}

	fmt.Println("File written successfully")
}

func ReadFile() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File contents:")
	fmt.Println(string(data))
}

type User struct {
	Name string `json:"name"`
	Age int     `json:"age"`
	Email string `json:"email"`
}

func EncodeJson() {
	user := User{Name: "Lukman", Age: 25, Email: "lukman@example.com"}

	//convert struct to JSON
	data, err := json.MarshalIndent(user, "", " ") // pretty print
	if err != nil {
		panic(err)
	}

	//save JSON to file
	err = os.WriteFile("user.json", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("User saved to user.json")
}

func DecodeJson() {
	data, err := os.ReadFile("user.json")
	if err != nil {
		panic(err)
	}

	var user User
	err = json.Unmarshal(data, &user) // decode into struct
	if err != nil {
		panic(err)
	}

	fmt.Println("Decode User Struct:")
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Email:", user.Email)
}

func JsonArray() {
	users := []User{
		{Name: "Lukman", Age: 25, Email: "lukman@example.com"},
		{Name: "Ada", Age: 30, Email: "ada@example.com"},
	}

	// encode list of users
	data, _ :=  json.MarshalIndent(users, "", " ")
	fmt.Println(string(data))

	// decode JSON back into slice
	var decoded []User
	json.Unmarshal(data, &decoded)

	fmt.Println("\nDecoded Users:")
	for _, u := range decoded {
		fmt.Printf("%s (%d) - %s\n", u.Name, u.Age, u.Email)
	}
}

func Calculator() {
	var a, b float64
	var op string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&op)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	switch op {
	case "+":
		fmt.Printf("Result: %.2f\n", a+b)
	case "-":
		fmt.Printf("Result: %.2f\n", a-b)
	case "*":
		fmt.Printf("Result: %.2f\n", a*b)
	case "/":
		if b != 0 {
			fmt.Printf("Result: %.2f\n", a/b)
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Invalid operator")
	}
}

func handleHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Hello, Go Web Server! 🚀")
	user := User{Name: "Lukman", Age: 25, Email: "lukman@example.com"}

	// set response header to JSON
	w.Header().Set("Content-Type", "application/json")

	// encode struct as JSON and send to client
	json.NewEncoder(w).Encode(user)
}

func WebServer() {
	http.HandleFunc("/user", handleHandler) // register route
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil) // start server
}

