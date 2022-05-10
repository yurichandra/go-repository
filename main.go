package main

import (
	"fmt"
	"sync"
	"time"
)

var clubIds = []int{1, 2, 3}

func main() {
	withoutConcurrency()

	concurrencyWithWaitGroupV1()

	consumeClosedChan()
}

func withoutConcurrency() {
	start := time.Now()

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

func concurrencyWithWaitGroupV1() {
	start := time.Now()

	wg := sync.WaitGroup{}
	clubChan := make(chan Club)
	doneChan := make(chan bool)
	errChan := make(chan error)
	count := 0

	for _, clubId := range clubIds {
		wg.Add(1)

		count++
		go func(wg *sync.WaitGroup, clubId int) {
			defer wg.Done()

			club, err := Fetch(clubId)
			if err != nil {
				errChan <- err
			}

			clubChan <- club
		}(&wg, clubId)
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()

		if count == len(clubIds) {
			close(doneChan)
		}
	}(&wg)

	clubs := make([]Club, 0)
	errs := make([]error, 0)

	for {
		select {
		case club := <-clubChan:
			clubs = append(clubs, club)
		case err := <-errChan:
			errs = append(errs, err)
		case <-doneChan:
			if len(errs) > 0 {
				fmt.Println(errs)
				return
			}

			elapsed := time.Since(start)

			fmt.Println("Time elapsed with concurrency wait group v1")
			fmt.Println(elapsed)
			fmt.Println(clubs)

			return
		}
	}
}
