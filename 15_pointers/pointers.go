package main

import "fmt"

func changeNum(num* int){
	*num = 5
	fmt.Println("in changenum", *num)
}

func main() {
	num := 1
	changeNum(&num)
	fmt.Println("after change num", num)
}
