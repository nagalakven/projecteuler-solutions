package main

import "fmt"

func main() {
	n := 1000
	cache := map[int]int{
		1: 3, 2: 3, 3: 5, 4: 4, 5: 4, 6: 3, 7: 5, 8: 5, 9: 4, 10: 3, 0: 0,
		11: 6, 12: 6, 13: 8, 14: 8, 15: 7, 16: 7, 17: 9, 18: 8, 19: 8, 20: 6,
		30: 6, 40: 5, 50: 5, 60: 5, 70: 7, 80: 6, 90: 6, 100: 7, 1000: 8,
	}

	letters := 0

	for i := 1; i <= n; i++ {

		if lenVal, exists := cache[i]; exists {
			letters += lenVal

			if i == 100 || i == 1000 {
				letters += cache[1] //for word "one"
			}
		} else {
			wordCnt := 0

			num := i

			//process 1000s
			if num >= 1000 {
				thousands := num / 1000
				roundUp := thousands * 1000

				//add "one"
				if thousands == 1 {
					wordCnt += cache[1]
				}

				if foundLen, exists := cache[roundUp]; exists {
					wordCnt += foundLen
				} else {
					length := cache[thousands] + cache[1000]
					wordCnt += length

					//update cache
					cache[roundUp] = length
				}

				//update num
				num %= 1000
			}

			//process 100s
			if num >= 100 {
				hundreds := num / 100
				roundUp := hundreds * 100

				//add "one"
				if hundreds == 1 {
					wordCnt += cache[1]
				}

				if foundLen, exists := cache[roundUp]; exists {
					wordCnt += foundLen
				} else {
					length := cache[hundreds] + cache[100]
					wordCnt += length

					//update cache
					cache[roundUp] = length
				}

				//update num
				num %= 100

				//add "and" if required
				if num > 0 {
					wordCnt += 3
				}
			}

			//process 10s, 1s
			if num > 0 {
				if foundLen, exists := cache[num]; exists {
					wordCnt += foundLen
				} else {
					tensRoundUp := (num / 10) * 10
					ones := num % 10
					length := cache[tensRoundUp] + cache[ones]
					wordCnt += length

					//update cache
					cache[num] = length
				}
			}

			cache[i] = wordCnt
			letters += wordCnt
		}
	}

	fmt.Println("Total letters count:", letters)
}
