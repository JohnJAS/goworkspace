package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var best int = math.MaxInt
	update := func(v int) {
		if math.Abs(float64(target-best)) > math.Abs(float64(target-v)) {
			best = v
		}
	}

	for i := 0; i < len(nums)-2; {
		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < target {
				//move left to make sum increase
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				update(sum)
			} else if sum > target {
				//move right to make sum decrease
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
				update(sum)
			} else {
				//curTarget == sum
				return target
			}
		}
		//去重
		i++
		for i < len(nums)-2 && nums[i] == nums[i-1] {
			i++
		}
	}

	return best
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
}
