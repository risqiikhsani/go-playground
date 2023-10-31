package slice

import (
	"fmt"

	"github.com/risqiikhsani/testplayground/array"
)

func TestSlice() {
	var s []string = array.A[1:3]

	// A map maps keys to values.
	var m = map[string]string{
		"a": "kucing",
		"b": "kucingimut",
	}

	fmt.Println(s)
	fmt.Println(m)
}
