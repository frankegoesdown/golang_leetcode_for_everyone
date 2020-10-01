package main

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for idx, i := range nums {
		if j, ok := index[target-i]; ok {
			return []int{j, idx}
		}
		index[i] = idx
	}
	return nil
}
