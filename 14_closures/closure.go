package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	//counter returns a function so increment is function

	//outer scope is not destroyed after being called here count variable

	fmt.Println(increment())

}
