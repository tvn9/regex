package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Compile the expression once
	var userID = regexp.MustCompile(`^[a-zA-Z]+\[[0-9]+\]$`)

	userIDArr := []string{"ThanhNg[51]", "KimNg[50]", "Mike{51}", "Tim55", "Marco(56"}

	for i, userid := range userIDArr {
		if userID.MatchString(userid) {
			fmt.Printf("#%d, Correct match %s\n", i, userid)
		} else {
			fmt.Printf("#%d, %s did not match\n", i, userid)
		}
	}
}
