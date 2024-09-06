package main

import (
	"fmt"
)

func returnLetterOcurrences(text string) map[string]int {
	letterOcurrences := make(map[string]int)
	for _, letter := range text {
		if _, exist := letterOcurrences[string(letter)]; exist {
			letterOcurrences[string(letter)] += 1
		} else {
			letterOcurrences[string(letter)] = 1
		}
	}
	return letterOcurrences

}

func main() {
	var text string
	fmt.Print("Enter the string: ")
	fmt.Scanln(&text)

	fmt.Println(returnLetterOcurrences(text))

}
