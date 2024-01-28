package input

import (
	"fmt"
)

func Test() {
	fmt.Println("Enter something :")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
