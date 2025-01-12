package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	rows := 0
	isDiag := false
	res := ""
	dict := make(map[int]string)

	if len(s) <= 1 || numRows <= 1 {
		return s
	}

	for i:=0; i < len(s); i++ {
		dict[rows] += string(s[i])
		if rows == numRows-1 {
			isDiag = true
		} else if rows == 0 {
			isDiag = false
		}
		if isDiag {
			rows -= 1
		} else {
			rows += 1
		}
	}
	for i := range numRows {
		res += dict[i]
	}
	return res
}

func main()  {
	a := convert("PAYPALISHIRING", 4)
	fmt.Println(a)
}