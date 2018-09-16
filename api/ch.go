package main

import (
	"time"
	"fmt"
	"sync"
)

func main() {
	 count := make(chan int)
	 defer close(count)
	 lock := sync.WaitGroup{}
	for i:=0; i< 10; i++ {
		lock.Add(1)
		defer lock.Done()
		go counter(i, count)
	}
	 for v:= range count {
		 lock.Wait()
	 	fmt.Println(v)
	 }
	 time.Sleep(time.Second * 2)
}

func counter(i int, out chan int) {
	out <- i
}
