package twoSum

func twoSum(nums []int, target int) []int {
	check := make(map[int]int)
	// map int to its index
	for idx, val := range nums {
		key, ok := check[target-val]
		if ok {
			return []int{key, idx}
		} else {
			check[val] = idx
		}
	}
	return nil
}

// test
