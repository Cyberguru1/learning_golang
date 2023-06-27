///////////////////////////////////////
//        !!Hamza Saidu     !!       //
//        !!Cyb3rguru       !!       //
//        !!Pointers  in Go !!       //
//////////////////////////////////////

package main

import (
	"fmt"
)

func main() {
	// pointers in Go

	name := "CyberHashira"
	fmt.Println(&name) // prints address of name
	fmt.Println()
	// storing address

	var x int

	ptr := &x

	fmt.Printf("ptr is type %T with value %v and address %p\n", ptr, ptr, &ptr)

	// declaring a pointer

	var ptr1 *float64
	_ = ptr1

	// creating a pointer using new

	p := new(int)

	x = 100

	p = &x

	fmt.Printf("p is of type %T, and content %v with value %v\n", p, *p, p)
	fmt.Printf("address of x is %p\n", &x)

	// changing the value of x using pointer

	*p = 223

	fmt.Println("value of x is :", x, "and that of *p is :", *p)

}
