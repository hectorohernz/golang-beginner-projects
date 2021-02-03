func removeElement(nums []int, val int) int {
	q := 0
	i := 0
	j := len(nums) - 1
	for j >= 0 && i <= j {
		if nums[i] == val {
			if nums[j] == val {
				nums[j] = 0
				q++
				j--
				continue
			} else {
				nums[i] = nums[j]
				nums[j] = 0
				q++
				j--
			}

		}
		i++
	}

	return len(nums) - q
}