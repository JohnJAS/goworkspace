package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		sumOfGas, sumOfCost, offset := 0, 0, 0

		for offset < len(gas) {
			j := (i + offset) % len(gas)
			sumOfGas += gas[j]
			sumOfCost += cost[j]

			if sumOfGas < sumOfCost {
				break
			}

			offset++

		}

		if offset == len(gas) {
			return i
		} else {
			i = i + offset
		}

	}
	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
}
