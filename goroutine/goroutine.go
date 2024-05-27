package goroutine_learning

import (
	"fmt"
	"sync"
)

func goroutineExample1() {
	go func() {
		fmt.Println("Hello from goroutineExample1")
	}()
}
func goroutineExample2() {
	go func(msg string) {
		fmt.Println(msg)
	}("Hello from goroutine Example2")
}
func goroutineExample3UsingWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello from goroutineExample3UsingWaitGroup 1")
		wg.Done()
	}()
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}("Hello from goroutine Example3UsingWaitGroup 2")
	wg.Wait()

}

func GoroutineInit() {
	fmt.Println("in main thread")
	//goroutineExample1()
	//goroutineExample2()
	goroutineExample3UsingWaitGroup()
}
