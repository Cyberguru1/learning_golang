///////////////////////////////////////
//        !!Hamza Saidu     !!       //
//        !!Cyb3rguru       !!       //
//        !!methods  in Go !!        //
//////////////////////////////////////

package main

import (
	"time"
	"fmt"
	"math"
)

type names []string 

func (n names) print(){ // emulation fo class in go

	for i, name := range n {
		fmt.Println(i, name)
	}
}

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrince int) {
	c.price = newPrince
	c.brand = newBrand
}

func (c car) changeCar(newBrand string, newPrince int){
	c.price = newPrince
	c.brand = newBrand
}

func (c *car) changeCar2(newBrand string, newPrince int){
	c.price = newPrince
	c.brand = newBrand
}

// interfaces in go simillar to class


// interface that links the functions
type shape interface {
	area() float64
	perimeter() float64
}

// creating the struct data type
type circle struct {
	raduis float64
}

type rectangle struct {
	width, height float64
}

// function methods of the classes

func area(r float64) float64 {
	return math.PI * r * 2
}

func perimeter(r float64) float64 {
	return 2 * (r.wdith + r.height)
}

// the print function
func print(s shape){
	fmt.Printf("\nShape: %#v\n", s)
	fmt.Printf("\nArea: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter()  )
}

func main() {

	// introduction to methods
	const day  = 24 * time.Hour

	fmt.Printf("%T\n", day)

	seconds := day.Seconds()

	fmt.Printf("%T\n", seconds)
	fmt.Println()
	fmt.Printf("Seconds in a day: %v\n", seconds)

	friends := names{"hamza", "Marry"} // first calling convention
	friends.print()

	fmt.Println()
	names.print(append(friends, "\n", "Sabo"))

	fmt.Println("integer method to time duration")
	var n int64 = 44854854895
	fmt.Println(n)
	fmt.Println(time.Duration(n))

	// methods with pointer
	fmt.Println("\nDealing with pointers in interfaces\n")
	myCar := car{brand:"Audi", price: 49999}
	changeCar(myCar, "Renault", 2000)
	fmt.Println("changing myCar features", myCar)

	myCar.changeCar("myCar", 21000)
	fmt.Println("changing myCar features", myCar)

	fmt.Println("\nUSing pointer:")
	myCar.changeCar2("myCarNaut", 21000)
	fmt.Println("changing myCar features", myCar)

	// using pointer
	
	var yourCar *car

	yourCar = &myCar

	fmt.Println("Using pointer to address:")
	yourCar.changeCar2("Hundait", 214400)
	fmt.Println("changing myCar features", *yourCar)

	//interfaces in go

	fmt.Println("Implementing interfaces in go\n")

	c1 := circle(raduis: 5.)
	r1 := rectangle(width:3., height: 2.1)

	print(c1)

	print(r1)
	

}
 