package main

import "fmt"

// func add(parameter) type {}
func add(a int, b int) int {
	return (a + b)
}
 //return many values
func getLanguages()(string, int, bool){
	return "golang", 12, true
}

func processIt(fn func(a int) int){
	fn(1)
}

func main() {
	result := add(3, 5)
	fmt.Println(result)
	fmt.Println(getLanguages())

	lang1 , lang2, lang3 := getLanguages();
	fmt.Println("lang1:", lang1, "lang2:", lang2, "lang3:", lang3)
	fn := func(a int)int {
		return 2
	}
	processIt(fn)
	
}
