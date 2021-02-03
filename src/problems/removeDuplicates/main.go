func removeDuplicates(nums []int) int {
	counter := 0
	i := 0
	for i < ((len(nums) - 1) - counter) {
		if nums[i] == nums[i+1] {
			for j := i + 1; j < len(nums); j++ {
				nums[j-1] = nums[j]
			}
			counter++
			continue
		}
		i++
	}
	return len(nums) - counter
}