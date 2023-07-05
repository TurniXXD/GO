package goroutines

import (
	"fmt"
	"time"
)

func bufferedChanCommunication() {
	// Communication via buffered chan is async and data is put into queue, goroutine blocks only after the limit of the queue is reached
	// FIFO Q (First in - First out queue) length is specified when initializing channel
	// Channel buffer has capacity of 3
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	// Get capacity of channel
	fmt.Println(cap(charChannel))

	// Get number of current occupied slots in channel
	fmt.Println(len(charChannel))

	charChannel <- "First write"
	charChannel <- "Second write"
	// Channel now has capacity of 2 occupied

	fmt.Println(len(charChannel))

	fmt.Println(<-charChannel)
	fmt.Println(<-charChannel)
	// Channel now has capacity of 0 occupied after its data is read

	fmt.Println(len(charChannel))

	for _, s := range chars {
		select {
		// After first value is extracted from chan the values will move and there will be two spots available in buffer
		case charChannel <- s:
		}
	}

	close(charChannel)

	fmt.Println(len(charChannel))

	// Residual data is still accessible in channel
	for result := range charChannel {
		fmt.Println(result)
	}

	// Residual data was extracted
	fmt.Println(len(charChannel))

}

func unbufferedChanCommunication(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("goroutine exited")
			return
		default:
			fmt.Println("main goroutine code")
			time.Sleep(time.Millisecond * 250)
		}
	}
}

func syncGoroutinesCommunication() {
	// Default chan capacity is 0
	done := make(chan bool)

	go unbufferedChanCommunication(done)

	time.Sleep(time.Millisecond * 1000)

	close(done)
}

// <-chan syntax means the return type is a READ ONLY channel
func sliceToChannel(nums []int) <-chan int {
	// With unbuffered channel goroutines communicate synchronously, they can only have single value
	out := make(chan int)
	// This goroutine doesn't block sliceToChannel
	go func() {
		for _, n := range nums {
			// goroutine blocks until the value is read from the channel
			out <- n
			fmt.Println("\nstage 1 chan read")
		}
		// Whe we close the channel it signals for loop in stage 2 to break the loop
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			time.Sleep(time.Millisecond * 500)
			// Stage 2 chan accepts data from stage 1 chan
			out <- n * n
			fmt.Printf("\nstage 2 channel read\n")
			// After it receives data it writes them into stage 2 channel and waits until there is new data on stage 1 channel
		}
		close(out)
	}()
	return out
}

func pipeline() {
	// input
	nums := []int{2, 3, 4, 7, 1}

	// stage 1
	dataChannel := sliceToChannel(nums)

	// stage 2
	finalChannel := sq(dataChannel)

	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}

// Goroutine is a independent path of execution
func Goroutines() {
	fmt.Println("\nGoroutines: ")

	bufferedChanCommunication()
	syncGoroutinesCommunication()
	pipeline()
}
