package utils

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func FindFactors(num int, factors map[int][]int) int {
	output := []int{}
	uniqueVals := make(map[int]struct{})

	if val, exists := factors[num]; exists {
		return len(val)
	}

	for divisor := 1; divisor*divisor <= num; divisor++ {
		if num%divisor == 0 {
			quotient := num / divisor

			//check if divisor is found in map
			if prevFactors, exists := factors[divisor]; exists {
				output = append(output, prevFactors...)
			} else {
				uniqueVals[divisor] = struct{}{}
			}

			//check quotient is found in map
			if quotient != divisor {
				if prevFactors, exists := factors[quotient]; exists {
					output = append(output, prevFactors...)
				} else {
					uniqueVals[quotient] = struct{}{}
				}
			}
		}
	}

	for _, val := range output {
		if _, exists := uniqueVals[val]; !exists {
			uniqueVals[val] = struct{}{}
		}
	}

	uniqueOutput := []int{}
	for key := range uniqueVals {
		uniqueOutput = append(uniqueOutput, key)
	}

	factors[num] = uniqueOutput

	return len(factors[num])
}
