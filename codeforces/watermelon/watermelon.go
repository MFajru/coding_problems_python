package main

import "fmt"

func  divideToTwo (weight int) string {
	if (weight > 2 && weight % 2 == 0) {
		return "YES"
	}
	return "NO"
}

func main() {
	var weight int;
	fmt.Scanln(&weight)
	fmt.Println(divideToTwo(weight))
}
