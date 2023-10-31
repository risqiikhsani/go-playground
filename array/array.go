package array

import (
	"fmt"
)

var A [5]string

func TestArray() {
	A[0] = "kucing"
	A[1] = "kelinci"
	A[2] = "ikan"

	fmt.Println(A[0], A[1])
	fmt.Println(A)

}
