package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := strconv.Atoi(os.Args[0])
	b := strconv.Atoi(os.Args[1])
	c := strconv.Atoi(os.Args[2])

	if a < 0 || b < 0 || c < 0 {
		fmt.Println("[!] All 3 argruments must be positive value.")
	} else if a > 200 || b > 200 || c > 200 {
		fmt.Println("[!] There is some value is out of range.")
	} else if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("[!] Not a triangle.")
	} else {
		aPow2 := a * a
		bPow2 := b * b
		cPow2 := c * c
		if aPow2+bPow2 == cPow2 || aPow2+cPow2 == bPow2 || bPow2+cPow2 == aPow2 {
			fmt.Println("[A] Right Triangle.")
		} else if a == b && b == c {
			fmt.Println("[A] Equilateral.")
		} else if a == b && a+b > c || a == c && a+c > b || b == c && b+c > a {
			fmt.Println("[A] Isosceles.")
		} else if a != b && a != c && b != c {
			fmt.Println("[A] Scalene.")
		}
	}
}
