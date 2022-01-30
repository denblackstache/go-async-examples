package main

import (
	"fmt"
	"time"
)

func myAsyncFunction(s int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		time.Sleep(time.Second)
		result <- s
	}()

	return result
}

func awaitSingleResult() {
	// aka async-await or Promise.resolve
	awaitResult := <-myAsyncFunction(1)
	fmt.Printf("await %d\n", awaitResult)
}

func awaitMultipleResults() {
	// aka Promise.all
	firstChannel, secondChannel := myAsyncFunction(2), myAsyncFunction(3)
	firstResult, secondResult := <-firstChannel, <-secondChannel
	fmt.Printf("await all %d,%d\n", firstResult, secondResult)
}

func awaitFirstAvailableResult() {
	// aka Promise.race
	var raceResult int
	select {
	case raceResult = <-myAsyncFunction(4):
	case raceResult = <-myAsyncFunction(5):
	}
	fmt.Printf("select/race %d\n", raceResult)
}

func main() {
	awaitSingleResult()         // 1 after 1s
	awaitMultipleResults()      // 2 3 after 1s
	awaitFirstAvailableResult() // 4 or 5 after 1s
}
