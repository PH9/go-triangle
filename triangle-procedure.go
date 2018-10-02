package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if validateArgs() == false {
		fmt.Println("Required 3 arguments")
		return
	}
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	c, _ := strconv.Atoi(os.Args[3])

	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("[!] All 3 arguments must be positive value.")
	} else if a > 200 || b > 200 || c > 200 {
		fmt.Println("[!] There is some value is out of range.")
	} else if a+b <= c && a+c <= b && b+c <= a {
		fmt.Println("[!] Not a triangle.")
	} else {
		aSqrt := a * a
		bSqrt := b * b
		cSqrt := c * c
		if aSqrt+bSqrt == cSqrt || aSqrt+cSqrt == bSqrt || bSqrt+cSqrt == aSqrt {
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

func validateArgs() bool {
	if len(os.Args) != 3 {
		return false
	}
	return true
}
