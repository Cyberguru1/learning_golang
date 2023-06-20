///////////////////////////////////////
//        !!Hamza Saidu  !!          //
//        !!Cyb3rguru    !!          //
//        !!Strings in Go!!          //
//////////////////////////////////////

package main


import (
	f "fmt"
	"strings"
)




func main () {

	print := f.Println
	printf := f.Printf
	_ , _= printf, print

	// declartion of strings in go

	var s1 string = "hello go world"

	s2 := "Hello golang!!"
	printf(" S2 is : %s and S3 is :%q\n",s2, s1)

	// concatenation in golang
	s3 := "I love " + "Golang" + "!!"

	print("after concatenation :: ")
	print(s3)

	// indexing in strings

	printf("Element at index 3 of s3: %v\n", string(s3[3]))

	// runes, bytes and unicodes

	var1 := 'a'

	printf("type of var1 is : %T with value : %v\n", var1, var1)

	// itreating through a string with for loop

	for _, i := range s3 {
		print(string(i))
	}

	for i:=0; i < len(s3); i++ {
		printf("%v", string(s3[i]))
	}

	// using the Strings module

	print("\n\n" + strings.Repeat("#", 12))

	// converting strings to runes

	name := "Hamza"

	rune_name := []rune(name)
	printf("\nThe name is %v, rune(int32) representations is %v with type %T\n", name, rune_name, rune_name)

	// using some of the strings module

	print("using the count function to count name: ",strings.Count(name, ""))

	// removing character using trim

	print(" triming ...hamza ?1424 to :", strings.Trim("...hamza 71424 ", ".123456789 "))

}