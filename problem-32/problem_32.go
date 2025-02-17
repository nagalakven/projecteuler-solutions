package main

import (
	"fmt"
	"strconv"
)

func generatePermutations(bytes []rune, start, end int, perms *[]string) {
	if start == len(bytes) {
		*perms = append(*perms, string(bytes))
	} else {
		for i := start; i < end; i++ {
			//fix 1 position and try all permutations
			bytes[i], bytes[start] = bytes[start], bytes[i]

			generatePermutations(bytes, start+1, end, perms)

			//revert to original
			bytes[i], bytes[start] = bytes[start], bytes[i]
		}
	}
}

func findPermutations(digits string) []string {
	var perms []string
	generatePermutations([]rune(digits), 0, len(digits), &perms)

	return perms
}

func calculatePandigitalProduct(perm string) int {
	pLen := len(perm)

	//multiplicand
	for i := 1; i <= pLen-2; i++ {
		//multiplier
		for j := i + 1; j < pLen-1; j++ {
			multiplicand, _ := strconv.Atoi(perm[:i])
			multiplier, _ := strconv.Atoi(perm[i:j])
			product, _ := strconv.Atoi(perm[j:])

			if (multiplicand * multiplier) == product {
				return product
			}
		}
	}

	return 0
}

func main() {
	digits := "123456789"
	permutations := findPermutations(digits)
	seenProducts := make(map[int]struct{})
	productsSum := 0

	for _, perm := range permutations {
		product := calculatePandigitalProduct(perm)
		if _, exists := seenProducts[product]; !exists {
			seenProducts[product] = struct{}{}
			productsSum += product
		}
	}

	fmt.Println("Pandigital products sum:", productsSum)
}
