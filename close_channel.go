package main

import (
	"fmt"
	"time"
)

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func closeChannel() chan int {
	numberChan := make(chan int)

	go func() {
		defer close(numberChan)

		for _, number := range numbers {
			numberChan <- number
			fmt.Println("writing to numberChan")
			time.Sleep(1 * time.Second)
		}
	}()

	return numberChan
}

func consumeClosedChan() {
	number := closeChannel()

	for {
		if val, ok := <-number; ok {
			println(val)
			fmt.Println("receiving from numberChan")
		} else {
			fmt.Println("exit reading from numberChan")
			return
		}
	}
}
