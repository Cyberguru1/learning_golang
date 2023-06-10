package main

import (
	"fmt"
	"strconv"
	"math"
)

func main() {
	var x int = 2 // using the var keyword

	var name string

	name = "hamza"

	var name2 = "Sadiu" // type inference

	//use bleak identifier {_ = name} to allocate unused variable

	xx := 3  // using the short operator works in block scopes ( // used only once )

	fmt.Println("second name:", name2)

	fmt.Println("first x and second xx with name", x, xx, name)

	// declaring multiple variables

	car, cost := "audi", 40000

	fmt.Println(car, cost, car)


	// Act of redeclaration with :=

	var opened = false

	opened, file := true, "a.txt" // used when a new variable is to be used

	_, _ = opened, file // not used in code again
	
	// declaring multiple variables

	var (
		salary float64
		firstname string
		gender bool
	)
	// by defalut amoounts to zero

	fmt.Println(salary, firstname, gender)


	// multiple declaration with 
	
	var a, b, c int

	fmt.Println(a, b, c)

	// swapping variables

	var i, j int

	i, j = 5, 6

	j, i = i, j // swapping variables

	fmt.Println(i, j)

	// suming tow varables

	sum := i + j 
    fmt.Println("summ is ", sum)
	sum = sum + 5 + 2

	fmt.Println("summ is ", sum)

	// type conversion and assignment

	fmt.Println("type conversion and assignment\n")

	var t int 

	t = 4

	var y float64

	y = 5.5

	fmt.Println(y)

	y = float64(t)

	fmt.Println(y)



	fmt.Println("zero values\n")
	// zero values: by default all varables are assigned to zero or none or false as in the case of string and bool resp.

	var names string

	var number int
	
	var coord float64

	var sex bool

	fmt.Printf("Contents are name: %s, number: %d, coord: %f, sex: %s\n",names, number, coord, sex)

	// comments

	fmt.Println("comments, inline comments\n")

	// example of inline comments

	fmt.Println("example of multplie line comment\n")

	/*

	multiple line comments

	*/
	// fmt package; standard library function

	fmt.Println("Dealing with fmt package")

	// using string formating

	fmt.Printf(" numbers in current workspace are %d, %s, %d, %f\n", x, name, y, t)

	// usnng escape sequence 

	fmt.Println(" HELLO GO WORLD \" HAMZA \" HERE !!!!\n")

	// more on percenting strings

	// %q to quote words

	// %v for all variable types

	// %x hexadecimal
	// %b binary 
	// %t bool type

	// 

	fmt.Printf("Helllo %v, %q\n", name, name)

	// %T -> type of variable also printed out

	fmt.Printf("my name is %T", name)

	// constants

	fmt.Println("Print dealing with constants\n")

	const pi float64 = 3.142

	fmt.Println(pi)
	const xxx = 5
	var (
		p int = xxx
		q float64 = xxx
		r byte = xxx
		bb = xxx > p
	)

	fmt.Printf("type of p is :%T\n type of q is :%T\n type of r is :%T\n", p, q, r)
	fmt.Printf("type of b is : %T with value %t\n",bb , bb)
	fmt.Println(p, q, r,"\n") 

	// iota constant

	fmt.Println("Dealing with iota constants:\n")
	const (
		C1 = iota // default 0
		C2 = iota
		C3
	)

	fmt.Println(C1, C2, C3)

	fmt.Println("manipulating with iota constants:\n")
	const (
		CC1 = iota * 2 // default 0
		CC2    // 2
		CC3
	)

	fmt.Println(CC1, CC2, CC3)

	fmt.Println("\n")
	const (
		aC1 = (iota * 2) + 1 // default 1
		aC2  // 3
		aC3
	)

	fmt.Println(aC1, aC2, aC3)

	const (
		m1 = -(iota + 2)  // -2
		_
		m2 
		m3
	)


	fmt.Println(m1, m2, m3)


	// Go Data types

	// int, array, slice,map, struct, pointer, fnction, interface, channel types

	// byte: alias for uint8
	// rune: alias for int32

	// complex64, complex128 : only the two types

	// examples
	// int8
	var i1 int8 = 120
	fmt.Printf("%T\n", i1)

	// unit16
	var i2 uint16 = 65530

	fmt.Printf("%T\n", i2)

	//rune TYPE

	var rr rune = 'A'
	fmt.Printf("%T, %v, %x\n", rr, rr,rr)

	//BOOL TYPE

	var bbb bool = true
	fmt.Printf("%T\n", bbb)

	//String TYPE

	var ss string = "HELLO GO!"
	fmt.Printf("%T\n, %ss\n", ss,ss)

	// Array and Slice Type
	
	var numbers = [4]int{1, 2, -3, 400}

	fmt.Printf("%T\n", numbers)

	//Slice Type :: no specified size

	var cities = []string{"london", "tokyo", "Germany"}
	fmt.Printf("%T\n", cities)

	// Map Type :: similar to dictionary 

	balances := map[string]float64{
		"USD" :223.2,
		"EUR": 533.2,
	}

	fmt.Printf("%T\n", balances)

	// Struct Type

	type Person struct{
		name string
		age int
	}

	var yo Person
	yo.name = "hamza"
	yo.age = 423

	fmt.Printf("%T, yo name {%v}, yo age {%v}\n", yo, yo.name, yo.age)

	// Pointer Type

	var xd int = 2
	ptr := &xd
	fmt.Printf("ptr is of type %T and value %v\n",ptr, ptr)
	
	// using pointer for struct

	pointer := &yo
	fmt.Printf("data type of pointer : %T, name {%v}, age {%v}\n",pointer, (*pointer).name, (*pointer).age)

	// function type

	fmt.Printf("ff is of type %T\n", ff)


	// overflow

	var xin int8 = 127

	fmt.Println(xin + 1, "\n")

	// type convertion

	y1 := "32421"

	y2 := fmt.Sprintf("%x", y1)
	y3 := fmt.Sprintf("%o", y1)
	fmt.Println(y2, y3, y1)

	// using strconv library

	var f1, _ = strconv.ParseFloat(y1, 64)
	fmt.Printf("Type of f1 is %T with value %f\n",f1, f1)

	var aa, _ = strconv.Atoi("50")
	fmt.Printf("type of aa is %T and value of %v\n", aa, aa)
	
	abb := strconv.Itoa(aa)
	fmt.Printf("type of bb is %T and value of %v\n", abb, abb)

	//Aliases

	//rune is an alias for int32 and byte for uint8

	fmt.Println("\nDealing with aliases types\n")

	var at byte = 10
	var bt uint8

	// since at and bt are of same type then

	bt = at
	fmt.Printf("'bt' is of type %T with value %v and 'at' is of value %v", bt, bt, at)
	
	fmt.Println("\n\nusing the type variable\n")

	type second = uint // used second as an alias for uint

	var hour second = 36000 // hour is of type uint

	fmt.Printf("hour is of type %T with value {%v}", hour, hour)

	const (
		jun int = iota + 6
		jul
		Aug
	)
	fmt.Println("\n","jun: ",jun, jul, Aug, "\n")
	
	noIpv6 := math.Pow(2, 128)

	fmt.Printf("type of reuslt %T, value is %b", noIpv6, noIpv6)
	
	fmt.Println("\n\n")
	const xfmt float64 = 1.422349587101

	fmt.Printf("xfmt in four decimal places is %.5v", xfmt)
}

func ff(){

}