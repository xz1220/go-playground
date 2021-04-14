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
