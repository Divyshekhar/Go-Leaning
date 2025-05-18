package main

import (
	"fmt"
	"maps"
)

func main(){
	//hash tables, objects, dictionary

	//creating maps
	// make(map[key]value)
	m := make(map[string]string)

	//setting an element

	m["name"] = "GoLang"
	m["area"] = "backend"

	//get an element

	fmt.Println(m)
	fmt.Println("checking for not existent")
	fmt.Println(m["random"]) //non existent then return zero value

	delete(m, "area")

	fmt.Println(m)
	clear(m)
	fmt.Println(m)

	t := map[string]int{"price": 40, "phone": 3}
	fmt.Println(t)

	//finding in map
	v, ok := t["price"] 
	// v = value
	// key = if exists then true else false

	if ok {
		fmt.Println("all ok", v)
	} else{
		fmt.Println("does not exist")
	}
	m2 := map[string]int{"price": 40, "phone": 3}

	fmt.Println(maps.Equal(m2, t))





}