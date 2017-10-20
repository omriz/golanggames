package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(input <-chan int, output chan<- int, index int) {
	fmt.Println("Hello from worker", index)
	output <- rand.Intn(10) + 1
	for {
		x := <-input
		fmt.Println("Worker", index, "got", x)
		output <- x * (rand.Intn(5) + 1)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(5001)+1))
	}
}

func main() {
	c12 := make(chan int, 1)
	c21 := make(chan int, 1)
	fmt.Println("Starting messaging.")
	go worker(c12, c21, 0)
	go worker(c21, c12, 1)
	var x int
	fmt.Scanln(&x)
}
