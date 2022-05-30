package algorithm

func TwoSum(nums []int, target int) [][]int{
	var res [][]int
	for k, v := range nums{
		for i := k + 1; i < len(nums); i++{
			if v + nums[i] == target {
				res = append(res, []int{k, i})
			}
		}
	}
	return res
}

func TwoSumHash(nums []int, target int) []int{
	hashMap := map[int]int{}
	for k, v := range nums{
		if s, ok := hashMap[target - v]; ok{
			return []int{s,k}
		}
		hashMap[v] = k
	}
	return nil
}