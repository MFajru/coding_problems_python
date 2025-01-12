package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func checkTeamVote(word1 string, word2 string, word3 string) int {
	coutOfOne := 0
	intWord1, err1 := strconv.Atoi(word1)
	intWord2, err2 := strconv.Atoi(word2)
	intWord3, err3 := strconv.Atoi(word3)
	if err1 == nil && err2 == nil && err3 == nil {
		if intWord1 == 1  {
			coutOfOne++
		}
		if intWord2 == 1 {
			coutOfOne++
		}
		if intWord3 == 1 {
			coutOfOne++
		}
	}
	

	if coutOfOne > 1 {
		return 1
	}
	
	return 0
}

func checkOne(strNum string) int {
	strings.TrimSpace(strNum)
	arrNum := strings.Split(strNum, " ")
	fmt.Printf("%T", strNum[0])
	if slices.Contains(arrNum, "1") {
		return 1
	}
	return 0
}

func main() {
	var numOfLines int
	// var word1 string
	count := 0
	fmt.Scanf("%d", &numOfLines)
	input := bufio.NewReader(os.Stdin)

	for i := 0; i < numOfLines; i++ {
		// fmt.Scanf("%s", &word1)
		word,_ := input.ReadString('\n')
		checkOne(word)

		// fmt.Println("word", word1)
		// count += checkOne(word1)
	}
	fmt.Println(count)
}