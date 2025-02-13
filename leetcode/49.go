package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	var r [][]string
	mp := make(map[string][]string)

	for _, str := range strs {
		sliceStr := []byte(str)
		sort.Slice(sliceStr, func(i, j int) bool {
			return sliceStr[i] < sliceStr[j]
		})
		sortedStr := string(sliceStr)
		mp[sortedStr] = append(mp[sortedStr], str)
	}

	for _, v := range mp {
		r = append(r, v)
	}

	return r
}

func main() {
	test := []byte("cba")
	sort.Slice(test, func(i, j int) bool {
		return test[i] < test[j]
	})
	fmt.Println(string(test))

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
