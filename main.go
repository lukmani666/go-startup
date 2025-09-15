package main

import (
	"fmt"
	"go-startup/basics"
)

func main() {
	basics.VariablesAndConstants()
	basics.DataTypes()
	basics.Operators()
	basics.InputOutput()
	basics.ControlFlow()
	basics.Loops()

	result1 := basics.Add(10, 20)
	fmt.Println("Sum:", result1)

	result2, message := basics.Divide(10, 2)
	fmt.Println("Result:", result2, "Message", message)
	
	_, errMsg := basics.Divide(5, 0)
	fmt.Println("ErrMsg:", errMsg) 

	//Example 1. No error
	if result3, err := basics.SafeDivide(10, 4); err == nil {
		fmt.Println("Result:", result3)
	} else {
		fmt.Println("Error:", err)
	}

	// Example 2. error case
	if result3, err := basics.SafeDivide(5, 0); err == nil {
		fmt.Println("Result:", result3)
	} else {
		fmt.Println("Error:", err)
	}

	basics.ArrayExample()

}