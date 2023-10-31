package testgoroutines

import (
	"fmt"
)

func Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func TestGoroutine2() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go Sum(s[:3], c)
	go Sum(s[3:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
