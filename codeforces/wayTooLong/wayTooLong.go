package main

import "fmt"

func wordAbbrv(strArr []string) {
	for i := 0; i < len(strArr); i++ {
		word := strArr[i]
		lenWord := len(word)
		lenBetween := lenWord - 2
		if (lenWord > 10) {
			fmt.Println(string(word[0]) + fmt.Sprint(lenBetween) + string(word[lenWord-1]))
			continue
		}
		fmt.Println(word)
	}
}

func main() {
	var numOfWord int
	var word string
	words := []string{}
	fmt.Scanln(&numOfWord)
	i := 0
	for i < numOfWord {
		fmt.Scanln(&word)
		words = append(words, word)
		i++
	}
	wordAbbrv(words)
}