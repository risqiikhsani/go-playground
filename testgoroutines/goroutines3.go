package testgoroutines

import (
	"fmt"
	"time"
)

func PrintMessage(a chan string) {
	time.Sleep(time.Second * 2)
	a <- "Hello from Goroutine"
}

func TestGoroutine3() {
	message := make(chan string)

	go PrintMessage(message)

	fmt.Println("Hello from main function")
	fmt.Println(<-message)
}
