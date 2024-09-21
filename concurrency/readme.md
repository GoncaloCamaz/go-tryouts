Concurrency overview
==============

*** Channels or Wait Groups ***

Wait groups are used for synchronization, channels are used to comunicate between go routines.

To use a wait group we should say how many go routines we are going to have. Then, in each go routing, we 
need to say that the routing is done with wg.Done()

So, we add counts to the wait group with

	wg.Add(<number of routines>)

Then, we call the routines

	go doStuff(&wg) -> inside this we call wg.Done() when its finish
	go doStuff(&wg) -> ...

Then we add wg wait to wait unitl everything is completed
	wg.Wait()


To communicate between go routines we cannot use returns, we must use channels

So, relating to channels, we have buffered and unbuffered channels.

A channel is always blocking if its full...
Therefore, a unbuffered channel is always full!! 

# Buffered channels

msgch := make(chan int, 10) // 10 being the number of ints the channel can accept

# Unbuffered channels --

msgch := make(chan int)

So, when we have a unbuffered channel, it will be a blocking channel. Therefore, in the code bellow
we will have a deadlock

func main() {
	msgch := make(chan int)

	msgch <- 10 // will block here...

	msg := <- msgch
	fmt.Println(msg)
}

To fix this we can create a go routing to write in the channel. Basically, we need to have something
ready to read from the channel

func main() {
	msgch := make(chan int)

	go func() {
		msg := <- msgch
		fmt.Println(msg)
	}

	msgch <- 10 // blocking here but the message is printed by the go routine
}

To use unbuffered channels, we need to be ready to read before start writting.

---