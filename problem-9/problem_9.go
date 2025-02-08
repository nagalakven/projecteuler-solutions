package main

import "fmt"

func main() {
	//a+b+c = 1000
	//a^2 + b^2 = c^2
	//a<b<c

	target := 1000

	aLimit := target / 3 //~limit value
	bLimit := target / 2 //~limit

	for a := 1; a < aLimit; a++ {
		for b := a + 1; b < bLimit; b++ {
			c := 1000 - a - b
			if (a*a)+(b*b) == (c * c) {
				res := a * b * c
				fmt.Printf("Pythagorean product[%d*%d*%d] : %d", a, b, c, res)
				return
			}
		}
	}
}
