package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("[S] Start application.\n[S] Converting args to int.")
	a, _ := strconv.Atoi(os.Args[0])
	b, _ := strconv.Atoi(os.Args[1])
	c, _ := strconv.Atoi(os.Args[2])
	fmt.Println("[S] Converted args to int.")

	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("[!] All 3 argruments must be positive value.")
	} else if a > 200 || b > 200 || c > 200 {
		fmt.Println("[!] There is some value is out of range.")
	} else if a+b <= c && a+c <= b && b+c <= a {
		fmt.Println("[!] Not a triangle.")
	} else {
		fmt.Println("[S] Values can posible be triangle.")

		fmt.Println("[S] Power 2 all values to prepare Right Triangle.")
		aPow2 := a * a
		bPow2 := b * b
		cPow2 := c * c
		fmt.Println("[S] Power 2 all values finished.")
		if aPow2+bPow2 == cPow2 || aPow2+cPow2 == bPow2 || bPow2+cPow2 == aPow2 {
			fmt.Println("[A] Right Triangle.")
		} else if a == b && b == c {
			fmt.Println("[A] Equilateral.")
		} else if a == b && a+b > c || a == c && a+c > b || b == c && b+c > a {
			fmt.Println("[A] Isosceles.")
		} else if a != b && a != c && b != c {
			fmt.Println("[A] Scalene.")
		}
		fmt.Println("[S] Finished check triangle type.")
	}
	fmt.Println("[S] End of application.")
}
