///////////////////////////////////////
//        !!Hamza Saidu     !!       //
//        !!Cyb3rguru       !!       //
//        !!functions in Go !!       //
//////////////////////////////////////

package main

import (
	"fmt"
	"strings"
	"math"
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
}

func varadic(a... int) {  // declaring a varadic function, args ends with 3 dots (...)
	fmt.Println("\nPrinting out all the input varables")
	for _,val  :=  range a {
		fmt.Println(val)
	}
}

func varadic1(a... int) {
	a[0] = 4000 // using varadic function to change value of first index
}

func sumAnDproduct(arr... int) (sum, prod int) {
	sum = 0
	prod = 1

	for _, val := range arr {
		sum += val
		prod *= val
	}

	return
}


func foo() {
	fmt.Println("This is foo()")
}


func foobarr() {
	fmt.Println("This is foobarr()")
}


func barr() {
	fmt.Println("This is barr()")
}

// function with params and varadic function

func parm_varadic(size int, names ...string) (answer string) {

	res  := strings.Join(names, "-*-")
	answer = fmt.Sprintf("Size : %v, Full Names: %v", size, res) 

	return
}


// anonymous function 

func increment(x int) func()  {
	return func() {
		x++
		fmt.Println(x)
	}
}



// cube function 


func cube(x float64) (result float64) {
	result = math.Pow(x, 3)
	return
}

func searchItem(item string, arr_str ...string) (bool) {
	fmt.Println(item)
	_ = result
	return
}

func main() {

	f1() // calling the function 
	
	// calling f2 and passing params

	f2(1, 2, 4, 5, "my name")

	// calling a function with a specified return type

	a, b := f3(4, 5)

	fmt.Println("called f3 with a and b",a, b )

	// returning without returning a variable in return statement

	fmt.Println("The sum of 1, 2, 4, is :", sum(1, 2, 4))

	// calling the varadic function

	varadic(1,3,4,5,5,6,6)

	// using list in varadic function

	fmt.Println("chainging the content of array using varadic function :")

	arr1 := []int{4, 5, 6, 6, 7}

	fmt.Println("Arr1 before call : ")

	fmt.Printf("%#v\n", arr1)

	varadic1(arr1...)

	fmt.Println("After call:: ")

	fmt.Printf("%#v\n", arr1)

	// finding some of an array using varadic function

	sum, prod := sumAnDproduct(arr1...)

	fmt.Printf("sum is : %v and product is : %v\n", sum, prod)

	// combining varadic function
	
	info := parm_varadic(445, "Wolfgang", "Amadeus", "Mozart")

	fmt.Println(info)

	// Defer statement : used to control call of function

	defer foo() // second fucntion to be called

	barr()

	defer foobarr() // first statement to be called just before exiting main

	// calling the annonymous function

	aa := increment(12)

	aa()
	aa()
	aa()

	// calling the cube function

	fmt.Println("the cube of 5.0 is :", cube(5.0))

	// searching for and item

	animals := []string{"lion", "tiger", "bear"}
	result := searchItem("bear", animals...)



	
}