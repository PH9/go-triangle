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
