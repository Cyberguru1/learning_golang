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
	radius float64
}


type rectangle struct {
	width, height float64
}


// function methods of the classes

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r rectangle) area() float64 {
	return  r.height * r.width
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

	c1 := circle{radius: 5.}
	r1 := rectangle{width:3., height: 2.1}

	v := c1

	fmt.Println(v.perimeter())

	fmt.Println("features of c1 :")
	print(c1)

	fmt.Println("\nfeatures of r1 :")
	print(r1)

	// empty type

	fmt.Println("\nEmpty data type using empty interfaces")
	var x interface{}

	x = 5

	fmt.Printf("x is %T and of type %#v\n", x, x)

	x = "my name is golang dev"
	
	fmt.Printf("x is %T and of type %#v\n", x, x)

	x = 43.2325

	fmt.Printf("x is %T and of type %#v\n", x, x)

}
 