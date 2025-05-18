package main

import "fmt"

func main() {
	//range is used for iterating over datastructures

	nums := []int{6, 7, 8}
	sum := 0
	// for index, num := range nums{}
	for _, num := range nums {
		fmt.Println(num)
		sum += num
		fmt.Println("sum = ", sum)

	}
	m := map[string]string{"fname": "john", "lname": "does"}

	for key, val := range m {
		fmt.Println(key, val)
	}
	//unicode code point rune
	//i is starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, c)
		fmt.Println(i, string(c))

	}

}
