package teststruct

import "fmt"

type Kucing struct {
	Name string
	Age  uint
}

func (c Kucing) Print() {
	fmt.Printf("Kucing name is %s and age is %d \n", c.Name, c.Age)
}

func (c *Kucing) ChangeName(s string) {
	c.Name = s
}
