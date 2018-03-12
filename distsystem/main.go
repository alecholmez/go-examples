package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Create 3 channels for each go routine
	hits1 := make(chan string, 1)
	hits2 := make(chan string, 1)
	hits3 := make(chan string, 1)

	// Create three go routines that will all ping the google server
	for i := 0; i < 50; i++ {
		go func() {
			time, err := HitGoogle()
			if err != nil {
				log.Fatal(err)
			}
			hits1 <- time
		}()
		go func() {
			time, err := HitGoogle()
			if err != nil {
				log.Fatal(err)
			}
			hits2 <- time
		}()
		go func() {
			time, err := HitGoogle()
			if err != nil {
				log.Fatal(err)
			}
			hits3 <- time
		}()

	HITS_LOOP:
		for {
			select {
			case time := <-hits1:
				fmt.Println("Channel 1: " + time)
				break HITS_LOOP

			case time := <-hits2:
				fmt.Println("Channel 2: " + time)
				break HITS_LOOP

			case time := <-hits3:
				fmt.Println("Channel 3: " + time)
				break HITS_LOOP

			}
		}

		time.Sleep(time.Second * 1)
	}

}

// HitGoogle will ping the google server once
func HitGoogle() (string, error) {
	// Create a stopwatch
	begin := time.Now()

	res, err := http.Get("http://google.com")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.Body == nil {
		return "", errors.New("Response body was nil")
	}
	if res.StatusCode != 200 {
		return "", errors.New("Expected response code 200")
	}

	// Get the end time for stopwatch
	end := time.Now()
	// Compare the two times to get result
	totalTime := end.Sub(begin)

	return totalTime.String(), err
}
