package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("[I] Please enter 3 parameters for 3 sides of triangle.")
		fmt.Println("    eg. " + os.Args[0] + " 3 4 5")
		return
	}

	var sides [3]int

	for i := 1; i < 4; i++ {
		side, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Println("[!] Please enter all 3 parameters with integer value.")
			return
		} else {
			sides[i-1] = side
		}
	}

	if sides[0] < 0 || sides[1] < 0 || sides[2] < 0 {
		fmt.Println("[!] All 3 argruments must be positive value.")
	} else if sides[0] > 200 || sides[1] > 200 || sides[2] > 200 {
		fmt.Println("[!] There is some value is out of range.")
	} else if sides[0]+sides[1] <= sides[2] || sides[0]+sides[2] <= sides[1] || sides[1]+sides[2] <= sides[0] {
		fmt.Println("[!] Not a triangle.")
	} else {
		aPow2 := sides[0] * sides[0]
		bPow2 := sides[1] * sides[1]
		cPow2 := sides[2] * sides[2]
		if aPow2+bPow2 == cPow2 || aPow2+cPow2 == bPow2 || bPow2+cPow2 == aPow2 {
			fmt.Println("[A] Right Triangle.")
		} else if sides[0] == sides[1] && sides[1] == sides[2] {
			fmt.Println("[A] Equilateral.")
		} else if sides[0] == sides[1] && sides[0]+sides[1] > sides[2] || sides[0] == sides[2] && sides[0]+sides[2] > sides[1] || sides[1] == sides[2] && sides[1]+sides[2] > sides[0] {
			fmt.Println("[A] Isosceles.")
		} else if sides[0] != sides[1] && sides[0] != sides[2] && sides[1] != sides[2] {
			fmt.Println("[A] Scalene.")
		}
	}
}
