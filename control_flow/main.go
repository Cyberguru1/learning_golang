///////////////////////////////////////
//        !!Hamza Saidu       !!     //
//        !!Cyb3rguru         !!     //
//        !!Control flow in Go!!     //
//////////////////////////////////////


package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	price, incstock := 100, true

	//basic if statement

	if price > 80 || incstock {
		fmt.Println("Too Expensive")
	}

	if price <= 100 && incstock {
		fmt.Println("Buy it")
	}
	// if-else statement

	fmt.Println("\nIF - ELSE statements")
	if price < 100 {
		fmt.Println("is cheap")
	} else if price == 100 {
		fmt.Println("IT's on edge")
	} else {
		fmt.Println("IT's Expensice")
	}

	// Command line arguments

	fmt.Println("\nCOmmand line args\n")

	fmt.Println("Os.Args", os.Args)

	fmt.Println("second elem:", os.Args[1])

	fmt.Println("Number of args", len(os.Args))

	var result, _ = strconv.ParseFloat(os.Args[1], 64)

	fmt.Println("second args is ", result)

	var res = fmt.Sprintf("%T", result)

	fmt.Println("converted result is: ", res)

	// simple if statements

	fmt.Println("\nsimeple if statements")

	i, err := strconv.Atoi("45")

	fmt.Println(err)

	if err == nil {
		fmt.Printf("type of i %T, and value %d", i, i)
	} else {
		fmt.Println("error occurred!!")
	}

	// simple statements

	if i, err := strconv.Atoi("43"); err == nil {
		fmt.Println("\nnew i variable :", i)
	} else {
		fmt.Println(err)
	}

	// For loops

	fmt.Println("\nFor loops: >>>>>>")
	for i := 0; i < 10; i++ {
		fmt.Println("current value of i is :", i)
	}

	// While loops

	fmt.Println("\n dealing with while loops")

	j := 10

	for j >= 0 {
		fmt.Println(" J now as value of : ", j)
		j--
	}

	sum := 0

	for sum <= 10 {
		sum++
	}
	fmt.Println("\n sum = ", sum)

	// for and continue statement

	fmt.Println("\n controling for loop with continue statement")

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("hit an odd number : ", i)
			continue
		}
		fmt.Println(i)
	}

	// for and break statement

	fmt.Println("\n controling for loop with break statement")

	count := 0

	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Println(i, " is divisible by 13")
			count++
		}
		if count == 10 {
			break
		}
	}
	fmt.Println("Done with current loop\n")

	// Label statement in loops

	fmt.Println("Dealing with label statement in loops\n")

	friends := []string{"Hamza", "demaria", "fuwa", "Yoshino"}
	names := [4]string{"Jurino", "Barba", "fuwa", "jamio"}
	// setting labels

outer:
	for index, friend := range friends {
		for _, name := range names {
			if friend == name {
				fmt.Println(name, " found in  friends at index", index)
				break outer
			}
		}
	}

	fmt.Println("\nAll done with current Label\n")

	// for loop using goto statement
	x := 0
loop:

	if x <= 10 {
		fmt.Println("value of x is ", x)
		x++
		goto loop
	}

	fmt.Println("\nDone with exection\n")

	// switch statement to emulate if statements
	fmt.Println("Switch  statements \n")

	language := "golang"

	switch language {

	case "Go":
		fmt.Println("Welcome to the go program!!")

	case "cmd", "python": // logical or
		fmt.Println("Hello PYthonista!!")

	case "go", "golang":
		fmt.Println("You've made it here!!\n")

	default:
		fmt.Println("Non of the options")
	}

	nn := 45

	switch true {

	case nn%2 == 0:
		fmt.Println(nn, "is even with numbers")

	case nn%2 != 0:
		fmt.Println(nn, "is odd with the numbers\n")

	default:
		fmt.Println("Non of the number options")

	}

	hour := time.Now().Hour()

	switch true {
	case hour < 12:
		fmt.Println("Go0d morining alaye!!!")
	case hour > 12:
		fmt.Println("Good afternoon/evening alaye!!!")
	default:
		fmt.Println("Alien time !!")
	}

	fmt.Println("All done with Scopes!!!!!!!")

}
