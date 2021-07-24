package main

import (
	"fmt"
	"strings"
)

// 比较两个字符串排序后是否一致

func compare(a string, b string) bool {
	astr := []rune(a)
	bstr := []rune(b)

	if len(astr) != len(bstr) {
		return false
	}

	for _, v := range astr {
		if strings.Count(a, string(v)) != strings.Count(b, string(v)) {
			return false
		}
	}
	return true
}

func main() {
	a, b := "nuhhkwer", "khwerhnu"
	bool := compare(a, b)
	fmt.Print(bool)
}
