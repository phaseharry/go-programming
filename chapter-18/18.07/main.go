package main

import "log"

func main() {
	s1 := 0
	output := make(chan int, 100)
	input := make(chan bool, 100)

	go push(1, 25, input, output)
	go push(26, 50, input, output)
	go push(51, 75, input, output)
	go push(76, 100, input, output)

	for c := 0; c < 100; c++ {
		input <- true
		num := <-output // blocks here until another goroutine sends something to the output channel to consume
		log.Println(num)
		s1 += num
	}
	log.Println(s1)
}

func push(from int, to int, input chan bool, output chan int) {
	for i := from; i <= to; i++ {
		<-input     // blocks here until there is a value within input channel that can be recieved from push function
		output <- i // once received, then we push out a value through output channel
	}
}
