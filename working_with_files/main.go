///////////////////////////////////////
//        !!Hamza Saidu  !!          //
//        !!Cyb3rguru    !!          //
//        !!Maps in Go   !!          //
//////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// Creating a file

	fmt.Println("Creating a new file")
	new_file, err := os.Create("a.txt") // returns a pointer to the newly created file

	if err != nil {
		// error exists
		log.Fatal(err)
	}

	_ = new_file

	// Truncating or empting a file

	err = os.Truncate("a.txt", 0) // o means empty the file

	if err != nil {
		log.Fatal(err)
	}

	defer new_file.Close()

	// openning a flie with more options

	new_file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)

	// handling errors
	if err != nil {
		log.Fatal(err)
	}

	new_file.Close()

	// writing to a file

	filewr, err := os.OpenFile(
		"a.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer filewr.Close() // called at the end of the function

	byteSlice := []byte("I learn Golang! ")
	fmt.Println(byteSlice)

	byteswritten, err := filewr.Write(byteSlice)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Byes written: %d\n", byteswritten)

	if ioutil.WriteFile("c.txt", []byte("My name is hamza\n name was defined"), 0644) != nil {
		log.Fatal(err)
	} // written and error handling

	// Reading in a file in go

	if data, er := ioutil.ReadFile("c.txt"); er == nil {
		fmt.Printf("%s", data)
	}

	// reading file using bufio

	file, err := os.Open("c.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// creating a scanner from file pointer
	fmt.Println()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // loop through and scan for new line \n
		fmt.Println(scanner.Text())
	}

	// Reading from stdin

	for result := bufio.NewScanner(os.Stdin); result.Scan(); {
		text := result.Text()
		fmt.Println("Yov've Entered: ", text)
		if text == "exit" {
			fmt.Println("Exitting ......")
			break
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
