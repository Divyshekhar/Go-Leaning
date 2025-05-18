package main

import (
	"fmt"
	"slices"
)

func main() {
	//dynamic arrays

	// unintialized slice is nil
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	//make([]type, size, initial_capacity)
	//capacity-> max numbers of elements can fit
	var nums = make([]int, 0, 5)

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)

	fmt.Println(cap(nums))

	fmt.Println(nums)

	//copy function

	var nums2 = make([]int, len(nums))
	copy(nums2, nums) //(destination, source)
	fmt.Println(nums, nums2)


	//slice operator

	var nums3 = []int{1, 2, 3}
	fmt.Println(nums3[:2])

	//slice package

	var nums4 = []int{1, 2} //dynamic array
	var nums5 = []int{1, 2}

	fmt.Println(slices.Equal(nums4, nums5))

	var nums2d = [][]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(nums2d)




}
