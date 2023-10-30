package testinterface2

import "fmt"

type Anjing struct {
	Name string
	Age  uint
}

func (c Anjing) Print() {
	fmt.Printf("Anjing name is %s and age is %d \n", c.Name, c.Age)
}

func (c *Anjing) ChangeName(s string) {
	c.Name = s
}
