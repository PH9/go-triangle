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

	if isInteger, _ := IsAllArgsAreInt(); !isInteger {
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

func IsAllArgsAreInt() (isTriangle bool, sides [3]int) {
	for i := 1; i < 4; i++ {
		side, err := strconv.Atoi(os.Args[i])
		if err != nil {
			isTriangle = false
			return
		} else {
			sides[i-1] = side
		}
	}
	isTriangle = true
	return
}
