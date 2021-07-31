// 替换字符串中的空格，要去字符串由大小写英文字母组成

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func replaceBlank(ts string) (s string, b bool) {
	if len([]rune(ts)) > 1000 {
		return ts, false
	}
	for _, v := range ts {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return ts, false
		}
	}
	return strings.Replace(ts, " ", "%20", -1), true
}

func main() {
	var aa = "abcedewnb uyh ewd"
	ns, _ := replaceBlank(aa)
	fmt.Printf(ns)
}
