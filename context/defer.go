package main

import (
	"fmt"
)

func Println() {
	fmt.Println("return without init ,return:", Demo())  // 打印结果为 return: 0
	fmt.Println("return without init ,return:", Demo2()) // 打印结果为 return: 2
}

func Demo() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}

func Demo2() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}
