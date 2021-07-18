/**
 * 利用channel 控制协程序
 * 使用两个协程交替打印数字与字母 如: 12AB34CD56EF...
 */

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	// 声明两个channel
	letter, number := make(chan int), make(chan int)
	syncWait := sync.WaitGroup{}
	str := ""
	// 先启动一个协程输出数字
	go func(s *string) {
		i := 1
		for _ = range number {
			// 数字转为字符串
			(*s) += strconv.Itoa(i)
			i++
			(*s) += strconv.Itoa(i)
			i++
			letter <- 1
		}
	}(&str)
	syncWait.Add(1)

	go func(wait *sync.WaitGroup, s *string) {
		i := 65 // ascii A
		for _ = range letter {
			if i > 90 {
				wait.Done()
				return
			}
			(*s) += string(i)
			i++
			(*s) += string(i)
			i++
			number <- 1
		}
	}(&syncWait, &str)

	// 先执行数字
	number <- 1
	// 等待执行完成
	syncWait.Wait()
	fmt.Printf(str)
}
