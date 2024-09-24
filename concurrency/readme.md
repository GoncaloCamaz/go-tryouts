Concurrency Overview
==============

# Channels or Wait Groups

Wait groups are used for synchronization, channels are used to communicate between go routines.

To use a wait group we should say how many go routines we are going to have. Then, in each go routine, we 
need to say that the routine is done with 
```
wg.Done()
```

So, we add counts to the wait group with
```
wg.Add(<number of routines>)
```
Then, we call the routines
```
go doStuff(&wg) -> inside this we call wg.Done() when its finish
go doStuff(&wg) -> ...
```
Then we add wg wait to wait until everything is completed
```
wg.Wait()
```

To communicate between go routines we cannot use returns, we must use channels

So, relating to channels, we have buffered and unbuffered channels.

A channel is always blocking if it's full...
Therefore, an unbuffered channel is always full!! 

# Buffered channels

We can use Buffered channels to do asynchronous communication since the go routine sending data to the channel will not block
by waiting until someone reads from the channel.
We can simply read the information later from the channel and proceed...

```
msgch := make(chan int, 10) // 10 being the number of ints the channel can accept
```
Considering the example below:

We want to send three chars through a channel

```
func main() {
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	// important to close the channel
	close(charChannel)

	// we can now loop through the result
	for result := range charChannel {
		fmt.Println(result)
	}
}
```

# Unbuffered channels

Unbuffered channels are used to do synchronous communication because, since there is no capacity, we must be ready to read when we send data.
If not, the go routine will block its execution.
```
msgch := make(chan int)
```
So, when we have an unbuffered channel, it will be a blocking channel. Therefore, in the code below
we will have a deadlock
```
func main() {
	msgch := make(chan int)

	msgch <- 10 // will block here...

	msg := <- msgch
	fmt.Println(msg)
}
```
To fix this we can create a go routing to write in the channel. Basically we need to have something
ready to read from the channel
```
func main() {
	msgch := make(chan int)

	go func() {
		msg := <- msgch
		fmt.Println(msg)
	}

	msgch <- 10 // blocking here but the message is printed by the go routine
}
```
To use unbuffered channels, we need to be ready to read before start writing.

---

# Important Examples

## Let's say we want the main routine to cancel children go routines that might be running without stopping...

We can create a done channel!
```
// please note that this is a read only channel here (<-chan)
func doWork(done <-chan bool) {
	for {
		select {
		case <- done:
			fmt.Println("Finishing work...")
			return
		default:
			fmt.Println("Doing some random work that is taking forever...")
		}
	}
}

func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	// When this done channel is closed, the doWork go routine will receive the signal and stop due to the return call.
	close(done)
}
```
## Let's implement a pipeline...

Let's imagine that we have a slice with numbers and we want to perform some operations on those numbers. We can create a "pipeline" and execute those
operations separately.

[1, 3, 6, 4] ---> stage 1 ---> stage 2 ---> ["1", "9", "36", "16"]

```
// Please note that the functions below are not blocking because we are returning the channels.
// This means that we are writing into a channel but we are passing it to the next go routine which will
// read from it
func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	
	go func() {
		for _, n := range in {
			out <- n * n
		}

		close(out)
	}()

	return out
}

func main() {
	nums := []int{2, 3, 4, 7, 1}

	// start pipeline
	//stage 1 - convert num slice into channel
	dataChannel := sliceToChannel(nums)

	//stage 2 - pass output from stage 1 to another stage
	finalChannel := sq(dataChannel)

	// stage 3 - end of the pipeline where we will print the result
	for _, n := range finalChannel {
		fmt.Println(n)
	}
}
```
## Let's implement a pipeline with a generator and some slow calculations in the middle of the pipeline..

Here we will generate random numbers and take a certain amount of numbers from that generator. Then, we will
check which numbers are prime numbers but, this operation is going to be a litle bit intensive for us to test some patterns.

To solve this issue, we can use some known patterns like fan-in and fan-out to try to speed up and do more calculations concurrently

Bellow we have the naive approach where we do not use this patterns.

```
// This will be our generator
func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <- done:
				return
			case stream <- fn(): 
			}
		}
	}()

	return stream
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <- chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
				case <-done:
					return
				// this means we are writing into taken
				case taken <- <-stream:
			}
		}
	}()

	return taken
}

// this function is really slow but its just to expose the problem...
func primeFinder(done <- chan int, randIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := randomInt -1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt := <-randIntStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()

	return primes
}

func main() {
	done := make(chan int)
	defer close(done)

	randomNumberFetcher := func() int { return rand.Intn(500000) }
	randIntStream := repeatFunc(done, randomNumberFetcher)

	primeStream := primeFinder(done, randIntStream)

	for rand := range take(done, primeStream, 10) {
		fmt.Println(rand)
	}
}
```

The problem above is that, the primeFinder is really slow... We could improve this by distributing the calculation by multiple
go routines...

How? using the fan-out to split the calculations into several go routines and then fan-in to merge the results into one channel.

Starting by the fan-out problem, we want to distribute the load into several channels. However, how much can we really improve the problem? It would be a good idea to know how many cpu's we have available 

```
func main() {
	done := make(chan int)
	defer close(done)

	randomNumberFetcher := func() int { return rand.Intn(500000) }
	randIntStream := repeatFunc(done, randomNumberFetcher)

	// fan out - split the calculations and distribute workload in the amount of cpu's we have available
	CPUCount := runtime.NumCPU()
	// we create an array with <CPUCount> channels
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i:= 0; i < CPUCount; i++ {
		// to each channel we assign the stream coming from primeFinder
		primeFinderChannels[i] = primeFinder(done, randIntStream)
	}
}
```

Now, let's merge everything back together...


```
func fanIn[T any](done <-chan int, channels ...<-chan T) <-chanT {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := rance c {
			select {
			case <- done:
				return
			case fannedIntStream <- i:
			}
		}
	}

	for _, c := channels {
		// we need to add a wait group to wait for each transfer to finish!
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}


func main() {
	done := make(chan int)
	defer close(done)

	randomNumberFetcher := func() int { return rand.Intn(500000) }
	randIntStream := repeatFunc(done, randomNumberFetcher)

	// fan out - split the calculations and distribute workload
	CPUCount := runtime.NumCPU()
	// we want to put on this channel the results of the primeFinder results
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i:= 0; i < CPUCount; i++ {
		primeFinderChannels[i] = primeFinder(done, randIntStream)
	}

	// fan in - merge all channels results into only one channel!
	fannedInStream := fanIn(done, primeFinderChannels...)

	for rand := range take(done, fannedInStream, 10) {
		fmt.Prinln(rand)
	}
}
```
