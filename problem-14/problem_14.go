package main

import "fmt"

func findCollatzSequenceLength(num int, collatz map[int]int) int {
	if num == 1 {
		return 1
	}

	if val, exists := collatz[num]; exists {
		return val
	}

	var next int
	if num%2 == 0 {
		next = num / 2
	} else {
		next = (3 * num) + 1
	}

	collatz[num] = 1 + findCollatzSequenceLength(next, collatz)
	return collatz[num]
}

func main() {
	collatz := make(map[int]int)
	max := 1000000
	longest, longestNum := 0, 1

	for i := 2; i < max; i++ {
		length := findCollatzSequenceLength(i, collatz)
		if length > longest {
			longest = length
			longestNum = i
		}
	}

	fmt.Println("Longest Collatz Sequence number:", longestNum)
}
