package main

import (
	"os"
	"runtime/trace"
)

func keepAllocate() {
	for i := 0; i < 10000; i++ {
		go func() {
			select {}
		}()
	}
}

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	keepAllocate()
}
