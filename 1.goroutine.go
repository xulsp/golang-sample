package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)

func requst(worker int, count *uint64) {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		panic(err)
	}
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	num := atomic.LoadUint64(count)
	err2 := ioutil.WriteFile("./"+strconv.FormatUint(num, 10)+".json", body, 0644)
	if err2 != nil {
		panic(err2)
	}
	fmt.Print(result["code"], worker, num, "\n")
}

func main() {
	var workers int = 5
	var count uint64 = 0
	for i := 0; i < workers; i++ {
		go func(n int) {
			for {
				requst(n, &count)
				atomic.AddUint64(&count, 1)
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Second * 100)
}
