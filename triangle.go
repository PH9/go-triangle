package main

import (
	"fmt"
	"os"
	"strconv"
)

const MAXIMUM_SIZE = 3
const MAXIMUM_ARGRUMENTS = MAXIMUM_SIZE + 1

func main() {
	if ShowHelp() {
		return
	}

	var sides [MAXIMUM_SIZE]int
	var isInteger bool
	if isInteger, sides = IsAllArgsAreInt(); !isInteger {
		fmt.Println("[!] Please enter all 3 parameters with integer value.")
		return
	}

	if IsAnyNegativeValue(sides) {
		fmt.Println("[!] All 3 argruments must be positive value.")
		return
	}

	if IsAnyValueMoreThanMaximumRange(sides, 200) {
		fmt.Println("[!] There is some value is out of range.")
		return
	}

	if IsEquilateral(sides) {
		fmt.Println("[A] Equilateral.")
	}
}

func ShowHelp() bool {
	if len(os.Args) != MAXIMUM_ARGRUMENTS {
		fmt.Println("[I] Please enter 3 parameters for 3 sides of triangle.")
		fmt.Println("    eg. " + os.Args[0] + " 3 4 5")
		return true
	}
	return false
}

func IsAllArgsAreInt() (isTriangle bool, sides [MAXIMUM_SIZE]int) {
	for i := 1; i < MAXIMUM_SIZE+1; i++ {
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

func IsAnyNegativeValue(sides [MAXIMUM_SIZE]int) bool {
	for _, sideWidth := range sides {
		if sideWidth < 0 {
			return true
		}
	}
	return false
}

func IsAnyValueMoreThanMaximumRange(sides [MAXIMUM_SIZE]int, maximumValueValidable int) bool {
	for _, sideWidth := range sides {
		if sideWidth > maximumValueValidable {
			return true
		}
	}
	return false
}

func IsEquilateral(sides [MAXIMUM_SIZE]int) bool {
	for i := 1; i < len(sides); i++ {
		if sides[0] != sides[i] {
			return false
		}
	}
	return true
}
