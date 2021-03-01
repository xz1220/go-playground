package main

import (
	"fmt"
)

// func main() {
// 	const task = 10
// 	message := make(chan int, task)
// 	var wg sync.WaitGroup
// 	wg.Add(task)

// 	for i := 0; i < task; i++ {
// 		message <- i
// 		go func(ch chan int) {
// 			defer wg.Done()
// 			fmt.Println(<-ch)
// 		}(message)
// 	}

// 	wg.Wait()

// }

// 自旋锁

// //原子操作的变量.
// var Count int32

// //实现一个自旋锁操作.
// func SpinLock(i int32, fn func()) {
// 	for { //一个死循环.
// 		//查看原子操作的值.如果相等则执行函数
// 		if n := atomic.LoadInt32(&Count); n == i {
// 			fn()
// 			atomic.AddInt32(&Count, 1) //然后原子操作自增.
// 			break                      //一定要跳出循环.
// 		}
// 		time.Sleep(time.Nanosecond)
// 	}
// }

// func main() {
// 	for i := int32(0); i < 10; i++ {
// 		go func(i int32) {
// 			fn := func() {
// 				fmt.Println(i)
// 			}
// 			SpinLock(i, fn)
// 		}(i)
// 	}
// 	//当原子操作自增count = 10就结束操作, 也可以睡一会儿
// 	SpinLock(10, func() {})
// }

// func

func main() {
	chs := [11]chan struct{}{}

	chs[0] = make(chan struct{})
	for i := 0; i < 10; i++ {
		chs[i+1] = make(chan struct{})
		go func(i int) {
			//从通道中取出元素，才往下执行，否则一直阻塞等待
			<-chs[i]
			defer close(chs[i])
			fmt.Println(i)

			chs[i+1] <- struct{}{}
		}(i)
	}
	//确保主协程能对chs[0]附上值
	chs[0] <- struct{}{}
	//取出通道中最后一个未被取出的值
	<-chs[10]

	return
}
