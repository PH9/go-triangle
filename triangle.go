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

	var sides [3]int
	var isInteger bool
	if isInteger, sides = IsAllArgsAreInt(); !isInteger {
		fmt.Println("[!] Please enter all 3 parameters with integer value.")
		return
	}

	if IsAnyNegativeValue(sides) {
		fmt.Println("[!] All 3 argruments must be positive value.")
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

func IsAnyNegativeValue(sides [3]int) bool {
	for _, sideWidth := range sides {
		if sideWidth < 0 {
			return true
		}
	}
	return false
}
