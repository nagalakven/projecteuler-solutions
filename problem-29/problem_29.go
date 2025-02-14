package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := 100
	seen := make(map[string]struct{})

	for a := 2; a <= n; a++ {
		for b := 2; b <= n; b++ {
			bigA := big.NewInt(int64(a))
			bigB := big.NewInt(int64(b))

			result := new(big.Int).Exp(bigA, bigB, nil)
			seen[result.String()] = struct{}{}
		}
	}

	fmt.Println("Distint powers count:", len(seen))
}
