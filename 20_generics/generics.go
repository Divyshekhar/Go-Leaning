package main

import "fmt"

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
func printSlice2[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}
//generic stack
type Stack[T any] struct{
	elements []T
}

func main() {
	nums := [] int{1, 2, 3}
	names := []string{"golang", "typescript"}
	printSlice(names)
	printSlice2(nums)

	myStack := Stack[int]{
		elements:  []int{1, 2, 3 },
	}
	fmt.Println(myStack)


}
