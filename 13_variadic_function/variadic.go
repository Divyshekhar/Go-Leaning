package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// func anyparameter(nums ...interface{})int{
// 	total := 0

// 	return total
// }

func main() {

	// pass n numbers of parameters = variadic function
	fmt.Println(1, 2, 3, 4, 5, "hello")
	nums := []int{1, 2, 3, 5}

	result := sum(10, 15, 12, 10)
	result2 := sum(nums...)

	fmt.Println(result, result2)
}
