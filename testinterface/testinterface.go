package testinterface

import (
	"github.com/risqiikhsani/testplayground/testinterface1"
	"github.com/risqiikhsani/testplayground/testinterface2"
)

func TestInterface() {
	var coba Animal
	coba = &testinterface1.Kucing{
		Name: "Test",
		Age:  2,
	}
	coba.Print()

	var coba2 Animal
	coba2 = &testinterface1.Kucing{
		Name: "Bob",
		Age:  2,
	}
	coba2.Print()

	var coba3 Animal
	coba3 = &testinterface2.Anjing{
		Name: "Jo",
		Age:  2,
	}
	coba3.Print()

	test1 := App{Something: coba3}
	test1.PrintChangenamePrint("ucok")
}
