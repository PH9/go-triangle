package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("[s] start application.")
	if len(os.Args) != 4 {
		fmt.Println("[I] Please enter 3 parameters for 3 sides of triangle.")
		fmt.Println("    eg. " + os.Args[0] + " 3 4 5")
		return
	}
	fmt.Println("[s] end check only 3 args.\n[s] init variable int[] to hold 3 args.")

	var sides [3]int
	fmt.Println("[s] init variable finished.")

	fmt.Println("[s] start loop to convert all args to int.")
	for i := 1; i < 4; i++ {
		fmt.Println("[s] args index at", i, "start.")
		side, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Println("[s] args index at", i, "is not an integer.")
			fmt.Println("[!] Please enter all 3 parameters with integer value.")
			return
		} else {
			fmt.Println("[s] args index at", i, "is an integer.")
			sides[i-1] = side
		}
	}
	fmt.Println("[s] end convert args to int loop.")

	fmt.Println("[s] start checking is triangle.")
	if sides[0] < 0 || sides[1] < 0 || sides[2] < 0 {
		fmt.Println("[!] All 3 argruments must be positive value.")
	} else if sides[0] > 200 || sides[1] > 200 || sides[2] > 200 {
		fmt.Println("[!] There is some value is out of range.")
	} else if sides[0]+sides[1] <= sides[2] || sides[0]+sides[2] <= sides[1] || sides[1]+sides[2] <= sides[0] {
		fmt.Println("[!] Not a triangle.")
	} else {
		fmt.Println("[s] Start check which tyoe of triangle is.")
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
		fmt.Println("[s] end check type")
	}
	fmt.Println("[s] end of application")
}
