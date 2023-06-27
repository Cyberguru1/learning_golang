///////////////////////////////////////
//        !!Hamza Saidu  !!          //
//        !!Cyb3rguru    !!          //
//        !!Maps in Go   !!          //
//////////////////////////////////////

package main

import (
	f "fmt"
	s "strings"
)

func main() {
	p := f.Println
	pf := f.Printf
	rep := s.Repeat

	_, _, _ = p, pf, rep

	// declaring maps in golang

	var employes map[string]string

	pf("Employee is of type %T, and value %v\n", employes, employes)

	//adding contents

	employe_details := map[string]string{} // cause can't directly assign

	employe_details["name"] = rep("#", 5)
	employe_details["contact"] = rep("#", 5)
	employe_details["email"] = rep("#", 9)

	pf("%#v\n", employe_details)

	//ideclaring map uising a map function

	new_map := make(map[int]string)

	new_map[0] = rep("@", 4)

	p(new_map)

	//declaring and initializing a map

	balances := map[string]float64{
		"bsd": 555.6,
		"USD": 343.22,
	}

	p(balances)

	// check if a value exists in an array

	val, res := balances["bsdd"]

	pf("value: in balance with key bsdd: %v, and result of operation: %v\n", val, res)

	val, res = balances["bsd"]

	pf("value: in balance with key bsd: %v, and result of operation: %v\n", val, res)

	// looping through a map

	for key, value := range balances {
		pf(" KEY : %q Value : %f\n", key, value)
	}

	// cloning a map in go

	friends := map[string]string{
		"first":  rep("@", 1),
		"second": rep("@", 2),
		"third":  rep("@", 3),
		"fourth": rep("@", 4),
		"fifth":  rep("@", 5),
	}

	p(friends)

	new_friends := make(map[string]string) // initialze a new map

	for key, value := range friends {
		new_friends[key] = value
	}

	pf("%#v\n", new_friends)

}
