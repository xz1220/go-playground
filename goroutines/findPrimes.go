package main

import "fmt"

func main() {
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)

	for num := 2; num <= 100; num++ {
		origin <- num
	}
	close(origin)
	<-wait
}

// 素数只与前面的素数的有关，若不能被前面的素数整除，那么这个数一定是素数
func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}
