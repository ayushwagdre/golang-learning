package main

import (
	channel_learn "ayush/channel"
	"fmt"
)

func main() {
	fmt.Println("Hello, World! in the Main")
	//interface_learn.InterfaceExample()
	//interface_learn.ChannelExample()
	//interface_learn.ChannelExample2()
	//interface_learn.Deadlock()
	//interface_learn.DeadlockResolve()
	//interface_learn.Deadlock2()
	//channel_learn.Deadlock2Resolve()
	//generic_learning.GenericLearning()
	//goroutine_learning.GoroutineInit()
	channel_learn.InitSelect()

}