package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	//i++是应为到offset无法到下一跳时，要向后一个
	for i := 0; i < length; i++ {
		totalGas, totalCost, offset := 0, 0, 0
		for offset < length {
			totalGas += gas[(i+offset)%length]
			totalCost += cost[(i+offset)%length]
			//当前的gas无法抵达下一跳时结束
			if totalGas < totalCost {
				break
			}
			offset++
		}
		if offset == length {
			return i
		} else {
			i += offset
		}
	}
	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
}
