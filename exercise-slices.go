package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// create slice
	sy := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		sx := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			sx[x] = uint8(x % (y + 1))
		}
		sy[y] = sx
	}
	return sy
}

func main() {
	pic.Show(Pic)
}
