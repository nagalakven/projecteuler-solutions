package main

import (
	"fmt"
)

func findCycleLength(num int) int {
	modVal := 1
	seen := make(map[int]int)

	for i := 0; ; i++ {
		modVal = (modVal * 10) % num

		if modVal == 0 {
			//cycle terminates
			return 0
		}

		if index, exists := seen[modVal]; exists {
			return i - index
		}

		seen[modVal] = i
	}
}

func main() {
	n := 1000
	maxD, maxLen := 0, 0

	for d := 2; d < n; d++ {
		//skip terminating numbers
		if d%2 == 0 || d%5 == 0 {
			continue
		}

		cycleLen := findCycleLength(d)
		if cycleLen > maxLen {
			maxLen = cycleLen
			maxD = d
		}
	}

	fmt.Println("d with longest decimel length:", maxD)
}
