package main

import "fmt"

func main(){

	//when size known before hand then we use array (fixed size)

	var nums [4]int

	// nums[0] = 1

	fmt.Println(nums) //prints the entire array just like that //p.s. Damn

	tums := [3]int{1,2,3} //declaration in single line
	fmt.Println(tums);



	nums2 := [2][2]int{{3,4},{5,6}}
	fmt.Println(nums2)
	

	//array length
	fmt.Println(len(nums))
}