package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	i := 0
	items := make([]int, 2)

	return func() int {
		defer func() { i++ }()

		if i > 1 {
			sum := items[0] + items[1]
			defer func() { items[i%2] = sum }()
			return sum
		} else {
			items[i%2] = i
			return i
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
