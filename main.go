/*
 * @Author: your name
 * @Date: 2021-09-07 09:28:42
 * @LastEditTime: 2021-09-18 16:02:51
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go-playground/main.go
 */
package main

import "fmt"

// 给定一个只包含大写英文字母的字符串S，要求你给出对S重新排列的所有不相同的排列数。
// 如：S为ABA，则不同的排列有ABA、AAB、BAA三种。

func findAllPattern(s string) int {

	// 构建map
	dic := make(map[rune]int)
	for _, r := range s {
		if num, ok := dic[r]; ok && num != 0 {
			dic[r] += 1
			continue
		}

		dic[r] = 1
	}

	// check
	for key, val := range dic {
		fmt.Println("key & val :", key, val)
	}

	// 尝试构建
	var ans int
	var dfs func()
	dfs = func() {
		// 终止条件
		var finished bool = true
		for key, val := range dic {
			if val != 0 {
				dic[key] -= 1
				dfs()
				dic[key] += 1

				finished = false
			}
		}

		if finished {
			ans += 1
		}
	}

	dfs()

	return ans
}

func main() {
	c := "ABCDEFGHHA"
	fmt.Println(findAllPattern(c))
}
