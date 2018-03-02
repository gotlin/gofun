package twosum

func twoSum(nums []int, target int) []int {

	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}

		}

	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	mapt := make(map[int]int)

	for i, v := range nums {

		val, ok := mapt[target-v]
		if ok {
			return []int{val, i}
		} else {
			mapt[v] = i
		}

	}
	return nil
}
