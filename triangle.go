package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if ShowHelp() {
		return
	}

	if !isAllArgsAreInt() {
		fmt.Println("[!] Please enter all 3 parameters with integer value.")
		return
	}
}

func ShowHelp() bool {
	if len(os.Args) != 4 {
		fmt.Println("[I] Please enter 3 parameters for 3 sides of triangle.")
		fmt.Println("  eg. " + os.Args[0] + " 3 4 5")
		return true
	}
	return false
}

func isAllArgsAreInt() bool {
	for i := 1; i < 4; i++ {
		if _, err := strconv.Atoi(os.Args[i]); err != nil {
			return false
		}
	}
	return true
}
