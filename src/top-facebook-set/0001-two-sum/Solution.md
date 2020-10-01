Solution
Approach 1: Brute Force

The brute force approach is simple. Loop through each element xxx and find if there is another value that equals to target−xtarget - xtarget−x.

```(golang)
func twoSum(nums []int, target int) []int {

	n := len(nums)

	for i, v := range nums {
		for j := i + 1; j < n; j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
```
Complexity Analysis

    Time complexity : O(n2)O(n^2)O(n2). For each element, we try to find its complement by looping through the rest of array which takes O(n)O(n)O(n) time. Therefore, the time complexity is O(n2)O(n^2)O(n2).

    Space complexity : O(1)O(1)O(1).

Approach 2: Two-pass Hash Table

To improve our run time complexity, we need a more efficient way to check if the complement exists in the array. If the complement exists, we need to look up its index. What is the best way to maintain a mapping of each element in the array to its index? A hash table.

We reduce the look up time from O(n)O(n)O(n) to O(1)O(1)O(1) by trading space for speed. A hash table is built exactly for this purpose, it supports fast look up in near constant time. I say "near" because if a collision occurred, a look up could degenerate to O(n)O(n)O(n) time. But look up in hash table should be amortized O(1)O(1)O(1) time as long as the hash function was chosen carefully.

A simple implementation uses two iterations. In the first iteration, we add each element's value and its index to the table. Then, in the second iteration we check if each element's complement (target−nums[i]target - nums[i]target−nums[i]) exists in the table. Beware that the complement must not be nums[i]nums[i]nums[i] itself!

```(golang)
func twoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))

	for i, v := range nums {
		sub := target - v
		if j, ok := m[sub]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}

	return nil
}
```


Complexity Analysis:

    Time complexity : O(n)O(n)O(n). We traverse the list containing nnn elements exactly twice. Since the hash table reduces the look up time to O(1)O(1)O(1), the time complexity is O(n)O(n)O(n).

    Space complexity : O(n)O(n)O(n). The extra space required depends on the number of items stored in the hash table, which stores exactly nnn elements.

Approach 3: One-pass Hash Table

It turns out we can do it in one-pass. While we iterate and inserting elements into the table, we also look back to check if current element's complement already exists in the table. If it exists, we have found a solution and return immediately.

```(golang)
func TwoSum3(nums []int, target int) []int {
	if nums == nil || len(nums) < 2 {
		return []int{-1, -1}
	}

	res := []int{-1, -1}

	intMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := intMap[target-nums[i]]; ok {
			res[0] = intMap[target-nums[i]]
			res[1] = i
			break
		}
		intMap[nums[i]] = i
	}
	return res
}
```

Complexity Analysis:

    Time complexity : O(n)O(n)O(n). We traverse the list containing nnn elements only once. Each look up in the table costs only O(1)O(1)O(1) time.

    Space complexity : O(n)O(n)O(n). The extra space required depends on the number of items stored in the hash table, which stores at most nnn elements.

