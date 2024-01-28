package main

import (
	"fmt"

	"github.com/risqiikhsani/testplayground/db"

	"github.com/risqiikhsani/testplayground/inputOutput"
)

// "github.com/risqiikhsani/testplayground/input"

//"runtime"

// "github.com/risqiikhsani/testplayground/array"
// "github.com/risqiikhsani/testplayground/slice"
// "github.com/risqiikhsani/testplayground/testgoroutines"
// "github.com/risqiikhsani/testplayground/testinterface"
// "github.com/risqiikhsani/testplayground/teststruct"
// "github.com/risqiikhsani/testplayground/testinterface"
// "github.com/risqiikhsani/testplayground/testinterface1"
// "github.com/risqiikhsani/testplayground/testinterface2"
//"github.com/risqiikhsani/testplayground/teststruct"

func main() {

	// teststruct.TestStruct()

	// testinterface.TestInterface()

	// testgoroutines.TestGoroutine()

	// testgoroutines.TestGoroutine4()

	// testgoroutines.TestGoroutine3()

	// testgoroutines.TestGoroutine2()

	// array.TestArray()

	// slice.TestSlice()

	// input.Test()

	db := db.Db()
	defer db.Close()

	next := true
	for next {
		fmt.Print("Get data (1) or insert data (2) or exit (3)? :")
		var input uint
		fmt.Scanln(&input)
		switch input {
		case 1:
			inputOutput.ShowData(db)
		case 2:
			fmt.Print("how many data will be inserted :")
			var amount uint
			fmt.Scanln(&amount, 3)
			for i := 0; i < int(amount); i++ {
				inputOutput.InputData(db)
			}
		case 3:
			next = false
			fmt.Println("Exited")
		default:
			fmt.Println("Invalid input")
		}
	}

}
