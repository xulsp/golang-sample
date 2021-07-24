package main

import (
	"fmt"
)

//  rune将字符串转化成 unicode 码点
// unicode要求所有字符以两到三个字节表示

func reverse(s string) (b bool, ss string) {
	str := []rune(s)
	length := len(str)
	fmt.Printf("长度:%d \n", length)
	if length > 5000 {
		return false, string(str)
	}
	for i := 0; i < length/2; i++ {
		str[i], str[length-1-i] = str[length-1-i], str[i]
	}
	fmt.Printf("%s \n", string(str))
	return true, string(str)
}

func main() {
	ss := "逆转字符串"
	reverse(ss)
}
