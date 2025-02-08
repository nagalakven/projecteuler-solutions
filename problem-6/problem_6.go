package main

import "fmt"

func main() {

	n := 100

	sum := (n * (n + 1)) / 2
	sqSum := (n * (n + 1) * ((2 * n) + 1)) / 6
	diff := (sum * sum) - sqSum

	fmt.Println("Sum square difference:", diff)
}
