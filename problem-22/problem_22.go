package main

import (
	"fmt"
	"main/utils"
)

func calculateWordScore(word string) int {
	score := 0

	for _, letter := range word {
		score += int(letter - 'A' + 1)
	}

	return score
}

func main() {
	filename := "0022_names.txt"
	namesList, err := utils.ReadFromFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	namesScore := 0
	for position, name := range namesList {
		namesScore += calculateWordScore(name) * (position + 1)
	}

	fmt.Println("Total name score:", namesScore)
}
