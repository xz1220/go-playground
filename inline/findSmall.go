package inline

func findSmall(nums []int, target int) int {

	// if len(nums) == 0 {
	// 	return -1
	// }

	// l, r := 0, len(nums)-1
	// for l < r {
	// 	middle := (l + r) / 2
	// 	if nums[middle] >= target {
	// 		r = middle - 1
	// 	} else {
	// 		l = middle + 1
	// 	}
	// }

	// if (l == len(nums)-1 && nums[l] != target) || (l+1 < len(nums) && nums[l] != target && nums[l+1] != target) {
	// 	return -1
	// } else if nums[l] == target {
	// 	return l
	// }

	// return l + 1
	// length := len(nums)
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
		// fmt.Println(l, r)
	}

	if nums[l] < target {
		return l + 1
	}
	// else if nums[(l+r)/2] == target {
	// 	for i := (l+r)/2 - 1; i >= 0; i-- {
	// 		if nums[i] == target {
	// 			ans -= 1
	// 		}
	// 	}
	// 	// return ans
	// }
	return l
}
