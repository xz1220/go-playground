package tencent

import (
	"fmt"
	"strconv"
)

var ans []string

func answer() {
	//
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ans = make([]string, 0)
	for i := 1; i < 9; i++ {
		// fmt.Println("-" + strconv.Itoa(cal(nums[len(nums)-i:])))
		test(nums[:len(nums)-i], 100-cal(nums[len(nums)-i:]), "+"+strconv.Itoa(cal(nums[len(nums)-i:])))
		test(nums[:len(nums)-i], 100+cal(nums[len(nums)-i:]), "-"+strconv.Itoa(cal(nums[len(nums)-i:])))
	}
}

func test(nums []int, target int, tempAns string) {
	if len(nums) <= 1 {
		if target == 1 {
			fmt.Println("1" + tempAns)
			ans = append(ans, "1"+tempAns)
		}
	}
	// fmt.Println(tempAns)

	for i := 1; i < len(nums); i++ {
		// fmt.Println("-" + strconv.Itoa(cal(nums[len(nums)-i:])))
		test(nums[:len(nums)-i], target-cal(nums[len(nums)-i:]), "+"+strconv.Itoa(cal(nums[len(nums)-i:]))+tempAns)
		test(nums[:len(nums)-i], target+cal(nums[len(nums)-i:]), "-"+strconv.Itoa(cal(nums[len(nums)-i:]))+tempAns)
	}
}

func cal(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans = ans*10 + nums[i]
	}
	return ans
}
