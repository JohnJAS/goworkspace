package main

import (
	"fmt"
	"math"
)

func mySqrt0(x int) int {
	if x == 0 {
		return 0
	}
	t := int(math.Exp(0.5 * math.Log(float64(x))))
	if (t+1)*(t+1) <= x {
		return t + 1
	}
	return t
}

func mySqrt(x int) int {
	l, r := 0, x
	var mid int
	for l <= r {
		mid = (l + r) / 2
		square := mid * mid
		if square == x {
			return mid
		}

		if square > x {
			r = mid - 1
		}

		if square < x {
			l = mid + 1
		}
	}
	return r
}

func main() {
	fmt.Printf("%d ", mySqrt0(1))
	fmt.Printf("%d ", mySqrt0(2))
	fmt.Printf("%d ", mySqrt0(3))
	fmt.Printf("%d ", mySqrt0(4))
	fmt.Printf("%d\n", mySqrt0(5))

	fmt.Printf("%d ", mySqrt(1))
	fmt.Printf("%d ", mySqrt(2))
	fmt.Printf("%d ", mySqrt(3))
	fmt.Printf("%d ", mySqrt(4))
	fmt.Printf("%d ", mySqrt(5))
}
