package main

import "fmt"

func main(){

	//need to assign value to a constant while declaration
	const name = "Golang"
	const age = 30

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)


}