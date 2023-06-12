package main


import (
	"fmt"

)

func main() {
	fmt.Println("Dealing with array\n")

	// declaring arrays in golang

	var ar1 = [4]int{1, 2, 3, 4} // fixed size array of length 4

	var ar2 = []float64{1.0, 3.3, 4.4, 2.3} // no fixed length, float64 array

	ar3 := [...]string{"Hamza", "Saidu", "Mugiwara", "Zoro"} // using short operator and spread for variable size

	fmt.Printf("%#v\n", ar1)
	fmt.Printf("%#v\n", ar2)
	fmt.Printf("%#v\n", ar3)

	// accessing array elements
	// using for loop
	fmt.Println("\nPrinting array elements of ar1 with for loop:\n")
	for index, value := range ar1 {
		fmt.Println("index :", index, "value :", value)
	}

	// using classic for loops
	fmt.Println("\nUising classical for loops to print content of ar3\n")
	for i:=0; i<len(ar3); i++ {
		fmt.Println("index :", i, "value :", ar3[i])
	}

	// changing the values of an array

	new_arr := [...]int{}




	
}