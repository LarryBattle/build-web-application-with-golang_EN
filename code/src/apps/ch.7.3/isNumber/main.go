// Checks to see if the passed argument is a number or not.
// Example run : `go run main.go -12.3`
package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Missing arguments");
		fmt.Println("Usage: program [string]")
		os.Exit(1)
	}
	numberPattern := `^[-+]?(\d+\.?)?\d+$`
	if m, _ := regexp.MatchString(numberPattern, os.Args[1]); m {
		fmt.Println("Number")
	} else {
		fmt.Println("Not number")
	}
}
