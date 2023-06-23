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
	fmt.Printf("a = %#v, b = %#v, c = %#v, d = %#v, e = %#v", a, b, c, d, e)
}

func main() {

	f1() // calling the function 


}