package main

import (

	//"runtime"
	"sync"

	"github.com/risqiikhsani/testplayground/testgoroutines"
	// "github.com/risqiikhsani/testplayground/testinterface"
	// "github.com/risqiikhsani/testplayground/testinterface1"
	// "github.com/risqiikhsani/testplayground/testinterface2"
	//"github.com/risqiikhsani/testplayground/teststruct"
)

var wg sync.WaitGroup

func main() {

	// z := &teststruct.Kucing{
	// 	Name: "Wi",
	// 	Age:  4,
	// }

	// z.Print()

	// c := teststruct.Kucing{
	// 	Name: "Bob",
	// 	Age:  2,
	// }

	// data := []teststruct.Kucing{
	// 	{
	// 		Name: "Bob",
	// 		Age:  1,
	// 	},
	// 	{
	// 		Name: "Kuy",
	// 		Age:  1,
	// 	},
	// 	{
	// 		Name: "Ze",
	// 		Age:  1,
	// 	},
	// }

	// c.Print()
	// for _, k := range data {
	// 	k.ChangeName("Ok")
	// 	k.Print()
	// }

	// var coba testinterface.Animal
	// coba = &testinterface1.Kucing{
	// 	Name: "Test",
	// 	Age:  2,
	// }
	// coba.Print()

	// var coba2 testinterface.Animal
	// coba2 = &testinterface1.Kucing{
	// 	Name: "Bob",
	// 	Age:  2,
	// }
	// coba2.Print()

	// var coba3 testinterface.Animal
	// coba3 = &testinterface2.Anjing{
	// 	Name: "Jo",
	// 	Age:  2,
	// }
	// coba3.Print()

	// test1 := testinterface.App{Something: coba3}
	// test1.PrintChangenamePrint("ucok")

	testgoroutines.TestGoroutine()

	// go woof()
	// meow()

	// // add one thing to wait for
	// wg.Add(1)
	// go woof()
	// meow()

	// // stop waiting when 0 things in waiting list
	// wg.Wait()

	testgoroutines.TestGoroutine4()

	testgoroutines.TestGoroutine3()

	testgoroutines.TestGoroutine2()
}
