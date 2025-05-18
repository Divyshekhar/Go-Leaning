package main

import "fmt"

func main() {
	// age := 15
	// if age >= 18 {
	// 	fmt.Println("Person is adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person is teenager")
	// } else {
	// 	fmt.Println("persin is a kid")
	// }
	var role = "admin"
	var hasPermission = false

	if role == "admin" || hasPermission {
		fmt.Println("yes")
	}
	if role == "admin" && hasPermission {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	//can define variable inside if construc
	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Person is teenager", age)
	}
	//go does not have ternary operator ( ? : )

}
