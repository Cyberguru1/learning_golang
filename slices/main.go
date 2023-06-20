package main

import (
	"fmt"
)

func main() {

	var cities []string

	fmt.Println("Slices in python, similar to array without length")
	fmt.Printf("Cities content is %v\n", cities)
	fmt.Printf("cities is NUll ?? %v\n", cities == nil)

	// initialinzing slices
	arr := [4]int{1, 2, 3, 4}
	arr2 := make([]int, 4)

	_ = arr2

	fmt.Println(arr)

	// arr = append(arr, 1) this an error only works with slice

	// creating a slice from type literals

	type states []string

	america := states{"maryland", "alabama", "virgina"}

	fmt.Println("States in america are:", america, "\n")

	// comparing slices

	fmt.Println("Comparing slices....\n")
	slice1 := []int{1, 1, 1, 1}

	slice2 := []int{2, 2, 2, 2}

	fmt.Printf("slice 1 is %#v, slice 2 is %#v\n", slice1, slice2)

	cmp1 := fmt.Sprintf("%v", slice1)

	cmp2 := fmt.Sprintf("%v", slice2)

	fmt.Printf("type of cmp1 is : %#v, and cmp2 is : %#v\n", cmp1, cmp2)

	var result bool = cmp1 == cmp2

	fmt.Printf("result of comparision is : %#v\n", result)

	slice1 = []int{2, 2, 2, 2}

	slice2 = []int{2, 2, 2, 2}

	fmt.Printf("\nslice 1 is %#v, slice 2 is %#v\n", slice1, slice2)

	cmp1 = fmt.Sprintf("%v", slice1)

	cmp2 = fmt.Sprintf("%v", slice2)

	fmt.Printf("type of cmp1 is : %#v, and cmp2 is : %#v\n", cmp1, cmp2)

	result = cmp1 == cmp2

	fmt.Printf("result of comparision is : %#v\n", result)

	// Backing array in Slices

	fmt.Println("\n Delaing with backing arrays\n")

	s1 := []int{10, 200, 40, 50}

	s2, s3 := s1[0:1], s1[1:3]

	fmt.Printf("s1: %#v, s2: %#v, s3: %#v\n", s1, s2, s3)

	s3[0] = 30

	fmt.Printf("s1: %#v, s2: %#v, s3: %#v\n", s1, s2, s3)

	// adding more items to a slice

	fmt.Println("\nAdding more items in a slice using 'append' method\n")

	nums := []int{}

	nums = append(nums, 4, 5, 6, 7)

	fmt.Printf("Current content of nums: %#v\n", nums)

	nums = append(nums, 98, 90)

	fmt.Printf("Current content of nums: %#v\n", nums)

	// printing the content of a slice using cap and len methods

	fmt.Println("Getting nums features with len and cap")
	fmt.Printf("Capacity of nums is : %#v and Length is : %#v\n", cap(nums), len(nums))

	// making a new slice from an existing slice

	new_slice := []int{3, 9, 11, 13}
	fmt.Println("New slice is : ", new_slice, "\n")

	dst_slice := make([]int, len(new_slice))

	fmt.Println("Transfering the content of the new_slice")

	// using the copy function to transfer the content of new_slice, returns len

	res := copy(dst_slice, new_slice)

	fmt.Printf("dst_slice is now : %#v and result of operation is %#v\n", dst_slice, res)

}
