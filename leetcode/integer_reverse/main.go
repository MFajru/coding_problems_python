package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {

	a := strconv.Itoa(x)
	temp := ""
	res := ""

	if a[0] == '-' {
		temp = string(a[0])
		a = a[1:]
	}

	for i := len(a); i > 0; i-- {
		res += string(a[i-1])
	}
	res = temp + res
	iRes,_ := strconv.Atoi(res)
	
	if iRes > 2147483647 || iRes < -2147483647 {
		return 0
	}
	
	return iRes
}

func reverse2 (x int) int {
	MAX := 2147483647
	isNeg := false
	res := 0

	if x < 0 {
		isNeg = !isNeg
		x *= -1
	}

	for x > 0 {
		res = (res * 10) + (x % 10)
		x /= 10
	}

	if res > MAX {
		return 0
	}
	
	if isNeg {
		return res * -1
	}

	return res
}

func main() {
	  t := reverse2(1534236469)
	  fmt.Println(t)
}