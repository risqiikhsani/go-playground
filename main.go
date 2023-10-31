package main

import (

	//"runtime"
	"sync"

	"github.com/risqiikhsani/testplayground/testgoroutines"
	"github.com/risqiikhsani/testplayground/testinterface"
	"github.com/risqiikhsani/testplayground/teststruct"
	// "github.com/risqiikhsani/testplayground/testinterface"
	// "github.com/risqiikhsani/testplayground/testinterface1"
	// "github.com/risqiikhsani/testplayground/testinterface2"
	//"github.com/risqiikhsani/testplayground/teststruct"
)

var wg sync.WaitGroup

func main() {

	teststruct.TestStruct()

	testinterface.TestInterface()

	testgoroutines.TestGoroutine()

	testgoroutines.TestGoroutine4()

	testgoroutines.TestGoroutine3()

	testgoroutines.TestGoroutine2()
}
