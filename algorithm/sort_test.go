package algorithm

import (
	"fmt"
	"testing"
)

func testCase() [][]int {
	return [][]int{
		[]int{1, 2, 3, 4, 5, 6, 6, 7, 8, 2, 3, 4},
		[]int{3, 4, 5, 6, 23, 2, 4},
	}
}

func TestSorting(t *testing.T) {
	nums := []int{1, 3, 542, 12, 3, 34, 35, 12, 1, 4, 45, 6, 67, 2, 2, 6, 67, 87, 8}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestHeapSort(t *testing.T) {
	nums := []int{1, 3, 542, 12, 3, 4, 2, 1}
	HeapSort(nums)
	fmt.Println(nums)
}

func TestFind(t *testing.T) {
	nums := []int{1, 1, 1, 3}
	// target := 3
	// ans := findNum(nums, 1)
	// fmt.Println(ans)
	for i := 0; i <= 4; i++ {
		fmt.Println(rightAns(nums, i), findNum(nums, i))
		if rightAns(nums, i) == findNum(nums, i) {
			fmt.Println("success")
		}
	}

}
