package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// 因为main函数将在say("hello")返回后（始终打印5次结果）退出，继而整个进程退出，所以say("world")会出现至多打印5次的结果
func main() {
	go say("world")
	say("hello")

	ch := make(chan int, 1)
	ch <- 1
	ch <- 2
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
}