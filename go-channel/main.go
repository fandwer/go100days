package main

import (
	"fmt"
	"strconv"
	"time"
)

func say(s string) {
	for i := 0; i < 100; i++ {
		fmt.Println(s + "的" + strconv.Itoa(i))
		time.Sleep(100 * time.Millisecond)

	}
}
func main() {
	q := strconv.Quote("Hello, 世界")
	fmt.Println(q)
	go say("Hello")
	go say("kettiy")
	say("word")
}
