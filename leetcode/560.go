package main

import "fmt"

func subarraySum(nums []int, k int) int {
	//设 prefixSum[i] = [0,i]的数之和
	//则 [j,i]的数之和 = prefixSum[i] - prefixSum[j-1], 其中j<=i,
	//即 prefixSum[j-1] == prefixSum[i] - k, 如果 i的前缀和 - k 已经存在，说明存在元素和为k的子串
	//又因为 j<=i, 在遍历i求prefixSum[i]的过程中，prefixSum[j-1]的值已经生成过了，所以不存在重复计算

	// 创建一个map来保存前缀和出现的次数
	prefixSumCount := make(map[int]int)
	prefixSumCount[0] = 1 // 初始化前缀和为0出现1次
	prefixSum := 0        // 初始化前缀和
	count := 0            // 记录和为k的子数组个数

	for _, num := range nums {
		prefixSum = prefixSum + num // 更新前缀和
		// 检查是否存在prefixSum - k的前缀和
		if countK, found := prefixSumCount[prefixSum-k]; found {
			count = count + countK
		}
		// 将当前前缀和加入map中
		prefixSumCount[prefixSum]++
	}

	return count
}

func main() {
	fmt.Println(subarraySum([]int{1}, 0))
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
}
