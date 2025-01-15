package main

import "math/rand"

type RandomizedSet struct {
	nums []int
	//value: index
	index map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		[]int{},
		map[int]int{},
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.index[val]; ok {
		return false
	}
	rs.index[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	if i, ok := rs.index[val]; !ok {
		return false
	} else {
		last := len(rs.nums) - 1
		lastVal := rs.nums[last]
		//处理数据存储区
		rs.nums[i] = lastVal
		rs.nums = rs.nums[:last]
		//处理索引区
		rs.index[lastVal] = i
		delete(rs.index, val)
	}
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
