package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt1(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		//fmt.Println(fmt.Sprintf("%d -> %v", i, z))
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt1(2))
	fmt.Println(Sqrt1(-2))
}
