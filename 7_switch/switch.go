package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:  //optional
		fmt.Println("other")
	}


	//multiple condition switch 

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday: //sat or sunday
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its work dat")
	}

	//type switch
 //interface{} = type = any
	whoAmI := func(i interface{}){
		switch i.(type){ //returns type of the variable .(type)
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("Its a bool")
		default:
			fmt.Println("Other")
		}
	}

	whoAmI(10.2)

}
