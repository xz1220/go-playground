package algorithm

// func QuickSort(nums []int, l, r int) {
// 	if l >= r {
// 		return
// 	}
// 	base := nums[l]
// 	left, right := l+1, r
// 	for {
// 		for left < len(nums) && nums[left] <= base {
// 			left++
// 		}
// 		for right >= 0 && nums[right] > base {
// 			right--
// 		}
// 		if left >= right {
// 			break
// 		}
// 		nums[left], nums[right] = nums[right], nums[left]
// 	}
// 	nums[l], nums[right] = nums[right], nums[l]
// 	fmt.Println(nums)
// 	// fmt.Println(l, left, right, r)
// 	QuickSort(nums, l, right-1)
// 	QuickSort(nums, left, r)
// }

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	base := nums[l]
	left, right := l+1, r

	for {
		// when left = r+1, and the loop will break
		for left <= r && nums[left] <= base {
			left++
		}

		for right >= l && nums[right] > base {
			right--
		}

		if left > right {
			break
		}

		nums[left], nums[right] = nums[right], nums[left]
	}

	nums[l], nums[right] = nums[right], nums[l]
	quickSort(nums, l, right-1)
	quickSort(nums, left, r)
}

//nums := []int{1, 3, 542, 12, 3, 4, 2, 1}
func HeapSort(nums []int) {
	// build heap
	length := len(nums)
	for i := length / 2; i >= 0; i-- {
		heapfy(nums, i, length)
	}

	//
	for i := length - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapfy(nums, 0, i)
	}

}

func heapfy(nums []int, i, length int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < length && nums[left] > nums[largest] {
		largest = left
	}
	if right < length && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]
		heapfy(nums, largest, length)
	}

}

// 给定一个有序数组 ns 和元素 e，找到最小的 i 使得 ns[i] >= e，如果不存在，返回 ns.length

// [1, 1, 1, 3], 0 => 0
// [1, 1, 1, 3], 1 => 0
// [1, 1, 1, 3], 2 => 3
// [1, 1, 1, 3], 3 => 3
// [1, 1, 1, 3], 4 => 4

func rightAns(nums []int, target int) int {
	var ans int
	for ans = 0; ans < len(nums) && nums[ans] < target; ans++ {
	}
	return ans
}

func findNum(nums []int, target int) int {
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
