///////////////////////////////////////
//        !!Hamza Saidu     !!       //
//        !!Cyb3rguru       !!       //
//        !!methods  in Go !!        //
//////////////////////////////////////

package main

import (
	"time"
	"fmt"
)

func main() {
	const day  = 24 * time.Hour

	fmt.Printf("%T\n", day)

	seconds := day.Seconds()

	fmt.Printf("%T\n", seconds)
	fmt.Println()
	fmt.Printf("Seconds in a day: %v\n", seconds)
}