package testgoroutines

import (
	"fmt"
	"time"
)

func SayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func TestGoroutine() {
	go SayHello()
	for i := 0; i < 5; i++ {
		fmt.Println("Hi")
		time.Sleep(100 * time.Millisecond)
	}

}
