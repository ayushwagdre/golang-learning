package channel_learning

import (
	"fmt"
	"sync"
)

//Channels in Go act as a medium for goroutines to communicate with each other.
//sometimes there might be situations where two or more goroutines need to communicate with one another.
//In such situations, we use channels that allow goroutines to communicate and share resources with each other.

//Create Channel in Go // make() function to create a channe
//channelName := make(chan int)

//Send data to the channel :- channelName <- data
//Receive data from the channel :- dataRecieved := <- channelName

//When a goroutine sends data into a channel, the operation is blocked until the data is received by another goroutine.
//When a goroutine receives data from a channel, the operation is blocked until another goroutine sends the data to the channel.

// Buffering in channels refers to the ability of a channel to hold a certain number of elements in a queue before it's received by a goroutine.
// When you create a buffered channel, you specify the buffer size as the second argument to make(chan type, bufferSize).
// Unbuffered Channels
// Unbuffered channels are synchronous channels where the sender goroutine blocks until the receiver goroutine receives the data.
// Buffered Channels
// Buffered channels are asynchronous channels where the sender goroutine does not wait for the receiver goroutine to receive the data.

func ChannelExample() {

	// create channel
	number := make(chan int, 3)
	message := make(chan string)
	ch := make(chan string)
	rData := make(chan string)

	// function call with goroutine
	go channelData(number, message)

	// retrieve channel data
	fmt.Println("Channel Data:", <-number)  //15
	fmt.Println("Channel Data:", <-message) //Learning Go channel

	number <- 15
	// Attempting to receive again from the same channel
	if value, ok := <-number; ok {
		fmt.Println("******Channel Data*****", value)
	} else {
		fmt.Println("******Channel Data*****", "No data")
	}

	//When a goroutine sends data into a channel, the operation is blocked until the data is received by another goroutine.
	// function call with goroutine
	go sendData(ch)
	// receive channel data
	fmt.Println(<-ch)

	//When a goroutine receives data from a channel, the operation is blocked until another goroutine sends the data to the channel.
	go receiveData(rData)
	fmt.Println("No data. Receive Operation Blocked")
	// send data to the channel
	rData <- "Data Received. Receive Operation Successful"
}

func channelData(number chan int, message chan string) {
	// send data into channel
	number <- 15
	message <- "Learning Go channel"

}

func sendData(ch chan string) {
	// data sent to the channel
	ch <- "Received. Send Operation Successful"
	fmt.Println("No receiver! Send Operation Blocked")

}

func receiveData(rData chan string) {
	// receive data from the channel
	fmt.Println(<-rData)

}

func ChannelExample2() {
	// create a channel
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	// send data to the channel
	go func() {
		defer wg.Done()
		fmt.Println("Sending data to the channel")
		for i := 0; i < 10; i++ {
			ch <- i
		}

	}()
	go func() {
		wg.Wait()
		fmt.Println("wating in  data to the channel")
		close(ch)
	}()

	for value := range ch {
		fmt.Println("Received:", value)
	}
}

func deadlockSendValues(myIntChannel chan int) {
	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}

}

func Deadlock() {
	myIntChannel := make(chan int)
	go deadlockSendValues(myIntChannel)
	//the main thread is waiting for the channel to be recieve but it will never recieve since loop is upto 5
	for i := 0; i < 6; i++ {
		fmt.Println(<-myIntChannel) //receiving value
	}
}

// by fefault channel are unbuffered there it will receive
// func Deadlock2() {
// 	myIntChannel := make(chan int)
// 	myIntChannel <- 11
// }

func deadlock2SendResolve(myIntChannel chan int) {
	myIntChannel <- 11
}

func Deadlock2Resolve() {
	myIntChannel := make(chan int)
	go deadlock2SendResolve(myIntChannel)
	fmt.Println("channel exmaple", <-myIntChannel)
}

func deadlockResolveSendValues(myIntChannel chan int) {
	for i := 0; i < 5; i++ {
		myIntChannel <- i //sending value
	}
	close(myIntChannel)
}

func DeadlockResolve() {
	myIntChannel := make(chan int)
	go deadlockResolveSendValues(myIntChannel)
	//the main thread is waiting for the channel to be recieve but it will never recieve since loop is upto 5
	for i := 0; i < 6; i++ {
		value, ok := <-myIntChannel //receiving value
		if ok {
			fmt.Println(value)
		} else {
			fmt.Println("Channel is closed")
			break
		}
	}
	for value := range myIntChannel {
		fmt.Println(value) //receiving value
	}
}
