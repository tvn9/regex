// Testing out standard library
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func MatchPhNumber(rwlns ...[]string) []string {
	phList := []string{}
	for _, lns := range rwlns {
		for _, ln := range lns {
			if len(ln) != 0 {
				match := regexp.MustCompile(`\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}`)
				str := match.FindString(ln)
				phList = append(phList, str)
			} else {
				continue
			}
		}
	}
	return phList
}

func main() {
	// The scanner is able to scan input by lines
	sc := bufio.NewScanner(os.Stdin)
	rwLines := []string{}
	for sc.Scan() {
		lns := sc.Text()
		rwLines = append(rwLines, lns)
	}

	phList := MatchPhNumber(rwLines)
	fmt.Println("Result line: ", phList, len(phList))
}
