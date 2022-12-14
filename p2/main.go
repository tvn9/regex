/*
Using the provided array, create a second array that only includes the numbers with the 801 area code. (The area code is the first 3 numbers.)
*/
package main

import (
	"fmt"
	"regexp"
)

func match(arr ...[]string) []string {
	// ex1 := `801-`
	// ex2 := `^801-`
	// ex3 := `801[\-.\s][0-9][0-9][0-9][\-.\s][0-9][0-9][0-9][0-9]` // this is bad form, but basic syntax is correct to match legal phone number in the usa
	// ex4 := `^801-\d\d\d-\d\d\d\d$` // Short hand \d
	ex5 := `801-\d{3}-\d{4}$` // short hand and repeate
	var newArr []string
	for _, phoneNums := range arr {
		if len(phoneNums) > 0 {
			match := regexp.MustCompile(ex5)
			for _, phNum := range phoneNums {
				if match.MatchString(phNum) {
					newArr = append(newArr, phNum)
				}
			}
		} else {
			fmt.Println("Found an empty phone listing")
		}
	}
	return newArr
}

func main() {
	phoneNums := []string{
		"801-766-9754", "801-545-5454", "435-666-1212", "801-796-8010", "435-555-9801", "801-009-0909", "435-222-8013", "801-777-665734"}
	empty := []string{"", "801-ABC-3442"} // Testing empty array

	fmt.Println(match(phoneNums, empty))
}
