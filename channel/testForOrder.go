package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	data := make(chan int)
	wg := sync.WaitGroup{}
	for index := 0; index < 3; index++ {
		wg.Add(1)
		if index == 0 {
			go func() {
				for i := 0; i < 5; i++ {
					<-data

				}
				wg.Done()
			}()
		} else {
			go func() {
				for i := 0; i < 5; i++ {
					Data := <-data
					fmt.Println(Data)
				}
				wg.Done()
			}()
		}

	}
	for i := 0; i < 15; i++ {
		data <- i
	}
	wg.Wait()
	return
}
