package main

import "fmt"

func fibonacciChannel(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}

	// 只有发送者才能关闭信道，而接收者不能。
	// 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacciChannel(cap(c), c)
	for i := range c{
		fmt.Println(i)
	}
}