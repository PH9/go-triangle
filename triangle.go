package main

import (
	"fmt"
	"os"
)

func main() {
	if ShowHelp() {
		return
	}
}

func ShowHelp() bool {
	if len(os.Args) != 4 {
		fmt.Println("Please enter 3 parameters for 3 sides of triangle.")
		fmt.Println("  eg. " + os.Args[0] + " 3 4 5")
		return true
	}
	return false
}
