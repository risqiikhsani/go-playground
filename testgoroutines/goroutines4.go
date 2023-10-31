package testgoroutines

import (
	"fmt"
	"sync"
)

func woof() {
	for i := 1; i <= 5; i++ {
		fmt.Println("woof ", i)
	}
	// remove one thing from waiting list
	wg.Done()
}

// func woof() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("woof ", i)
// 	}
// }

func meow() {
	for i := 1; i <= 5; i++ {
		fmt.Println("meow ", i)
	}
}

func bar() {
	for i := 1; i <= 5; i++ {
		fmt.Println("meow ", i)
	}
}

var wg sync.WaitGroup

func TestGoroutine4() {
	// go woof()
	// meow()

	// add one thing to wait for
	wg.Add(1)
	go woof()
	meow()

	// stop waiting when 0 things in waiting list
	wg.Wait()
}
