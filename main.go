package main

import (
	"fmt"
	"time"
)

func main() {
	withoutConcurrency()
}

func withoutConcurrency() {
	start := time.Now()
	clubIds := []int{1, 2, 3}

	clubs := make([]Club, 0)

	for _, clubId := range clubIds {
		club, err := Fetch(clubId)
		if err != nil {
			fmt.Println(err)
			break
		}

		clubs = append(clubs, club)
	}

	elapsed := time.Since(start)
	fmt.Println(clubs)

	fmt.Println("Time elapsed without concurrency")
	fmt.Println(elapsed)
}
