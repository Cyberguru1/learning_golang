///////////////////////////////////////
//        !!Hamza Saidu  !!          //
//        !!Cyb3rguru    !!          //
//        !!files in Go  !!          //
//////////////////////////////////////

package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	pf := fmt.Printf
	_, _ = p, pf

	// declaring structs

	type contact struct {
		name, email, address string
		id                   float64
	}

	pf("%#v\n", strings.Repeat("%", 4))

	details := contact{
		"hamza",
		"hamza@cyb3rguru.com",
		"127.0.0.1",
		99.9,
	}

	p(details)

	//accessing items

	pf("name : %#v\n", details.name)

	// updating a value

	details.name = "Shazam"

	pf("name : %#v\n", details)

	// initializing an anonymous stuct datatype

	Employee := struct {
		name, shape, size, height string
	}{
		name:   "ujiro",
		shape:  "##cm",
		size:   "###",
		height: "##",
	}

	pf("%#v\n", Employee)

}
