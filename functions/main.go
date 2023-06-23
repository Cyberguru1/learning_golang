///////////////////////////////////////
//        !!Hamza Saidu     !!       //
//        !!Cyb3rguru       !!       //
//        !!functions in Go !!       //
//////////////////////////////////////

package main

import (
	"fmt"
)


// defining function

func f1() {
	fmt.Println("F1 was called")
}

func f2(a, b, c, d int, e string) {
	fmt.Printf("a = %#v, b = %#v, c = %#v, d = %#v, e = %#v\n", a, b, c, d, e)
}

func f3(a int, b int) (int, int) {
	return a+b, b * a
}

func sum(a, b, c int) (s int) {
	s = a + b + c
	return // naked return statement

func main() {

	f1() // calling the function 
	
	// calling f2 and passing params

	f2(1, 2, 4, 5, "my name")

	// calling a function with a specified return type

	a, b := f3(4, 5)

	fmt.Println("called f3 with a and b",a, b )

	// returning without returning a variable in return statement

	fmt.Println("The sum of 1, 2, 4, is :", sum(1, 2, 4))

}