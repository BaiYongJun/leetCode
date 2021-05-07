package main

import "fmt"

func main() {
	aa := []int{}
	bb := []int{1, 3, 4, 5, 6, 7}
	aa = twoSum(bb, 6)
	fmt.Println(aa)
}

func twoSum(nums []int, target int) []int {
	globalmap := make(map[int]int)
	for index, num := range nums {
		targetNum := target - num
		i, ok := globalmap[targetNum]
		if ok {
			return []int{i, index}
		} else {
			globalmap[num] = index
		}
	}
	return []int{}
}
