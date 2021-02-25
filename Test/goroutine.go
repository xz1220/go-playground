package main

import "fmt"
import "sync"

func main() {
	const task = 10
	message := make(chan int, task)
	var wg sync.WaitGroup
	wg.Add(task)

	for i := 0; i < task; i++ {
		message <- i
		go func(ch chan int) {
			defer wg.Done()
			fmt.Println(<-ch)
		}(message)
	}

	wg.Wait()

}
