package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// We'll use the background context for the program but you can use whatever you want
	ctx, cancel := context.WithDeadline(context.Background(), get10())
	// Always call the cancellation function even if you won't ever hit that edge case
	defer cancel()

	// We'll make our channels here with a buffer of 1 since we don't need any WWWWRRRR. We're fine with WRWRWR
	done := make(chan bool)
	errs := make(chan error)

	// Call our DoSomething function so our program does something :)
	// We start this in a go routine to run it concurrently
	go DoSomething(ctx, done, errs)

	// Read over our done channel until we get true
	// Let's also check for errors in case we get any
	// If we exceed the context defined deadline it will send back on the errs channel
	for {
		select {
		case <-done:
			fmt.Println("Finished pinging google five times")
			return
		case e := <-errs:
			if e != nil {
				fmt.Println(e)
				return
			}
		}
	}
}

// getTime will calculate 10 seconds from now
func get10() time.Time {
	return time.Now().Add(time.Duration(10) * time.Second)
}

// DoSomething ...
func DoSomething(ctx context.Context, done chan bool, errs chan error) {
	var counter int

	for {
		select {
		case <-ctx.Done():
			// If we find that we've passed the deadline, return the context error
			errs <- ctx.Err()
		default:
			if counter == 5 {
				done <- true
			}

			fmt.Println("Pinging Google. . .")
			res, err := http.Get("http://google.com/")
			if err != nil {
				// Hopefully we never get this log case
				errs <- err
			}
			// Do nothing with the response body since we don't care
			defer res.Body.Close()

			time.Sleep(time.Second * 3)
			counter++
		}
	}
}
