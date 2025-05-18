package main

import "fmt"

func main() {
	//in go there is only for loop

	//while loop implementation using for
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//infinite loop
	// for{
	// 	println("1")
	// }

	//classic for loop

	// for i := 0; i < 3; i++ {
	// 	// break
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)

	// }

	//range
	for i:= range 3{
		fmt.Println(i)
	}

}
