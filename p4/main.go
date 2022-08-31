/*
The content.js file contains a string of text stored in a variable text1. This string of text is a statement that contains a day of the week as a part of the statement. Write a regular expression that will match any day of the week and then replace that day of the week in the string with Monday. Display the results to the console.
*/
package main

import (
	"fmt"
	"regexp"
)

func matchDay(str, repStr string) string {
	newString := str
	reg1 := `\b\w{1,4}?day\b`
	match := regexp.MustCompile(reg1)
	return match.ReplaceAllString(newString, repStr)
}

func main() {

	// monday, tuesday, wednesday, thursday, friday, saturday, sunday
	// find Tuesday, replace with Monday
	text := "Each and every Tuesday, at the beginning of the day, we hold a staff meeting. At the Tuesday staff meeting, you will need to make a report on the past weeks progress, and you will receive assignments for the following Tuesday. Just be aware that somedays this Tuesday meeting might not occur. When that happens, we will make an announcement."

	replace := "Monday"

	monday := matchDay(text, replace)

	fmt.Println("Before replacement")
	fmt.Println(text)
	fmt.Println()
	fmt.Println("After replaycement")
	fmt.Println(monday)
}
