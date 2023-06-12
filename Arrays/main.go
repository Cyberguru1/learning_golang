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
	fmt.Println("\nUsing classical for loops to print content of ar3\n")
	for i:=0; i<len(ar3); i++ {
		fmt.Println("index :", i, "value :", ar3[i])
	}

	// changing the values of an array

	new_arr := [5]int{}
	fmt.Printf("\nCurrent array content:%#v\n", new_arr)
	
	fmt.Println("changing the array content:\n")

	for index, _ := range new_arr {
		new_arr[index] = (index * 2) + 1
	}

	fmt.Printf("After changing : %#v\n", new_arr)

	// Declaring an array with new lines

	new_arr2 := [...]int{
		1,
		4,
		5,
		6,
		7, // must always end with coma
	}

	fmt.Printf("\n%#v\n", new_arr2)
	
	// keyed elements, simillar to dictionary

	key_arr := [...]int{
		1 : 5,
		0 : 44,
		2 : 55,
		4 : 20,
		3 : 12, // order doesen't matter, what matters for arrangement is the keys
	}

	fmt.Printf("\nKey Elements: %#v\n", key_arr)

}