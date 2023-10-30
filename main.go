package main

import (
	"github.com/risqiikhsani/testplayground/testinterface"
	"github.com/risqiikhsani/testplayground/testinterface1"
	"github.com/risqiikhsani/testplayground/testinterface2"
	"github.com/risqiikhsani/testplayground/teststruct"
)

func main() {

	z := &teststruct.Kucing{
		Name: "Wi",
		Age:  4,
	}

	z.Print()

	c := teststruct.Kucing{
		Name: "Bob",
		Age:  2,
	}

	data := []teststruct.Kucing{
		{
			Name: "Bob",
			Age:  1,
		},
		{
			Name: "Kuy",
			Age:  1,
		},
		{
			Name: "Ze",
			Age:  1,
		},
	}

	c.Print()
	for _, k := range data {
		k.ChangeName("Ok")
		k.Print()
	}

	var coba testinterface.Animal
	coba = &testinterface1.Kucing{
		Name: "Test",
		Age:  2,
	}
	coba.Print()

	var coba2 testinterface.Animal
	coba2 = &testinterface1.Kucing{
		Name: "Bob",
		Age:  2,
	}
	coba2.Print()

	var coba3 testinterface.Animal
	coba3 = &testinterface2.Anjing{
		Name: "Jo",
		Age:  2,
	}
	coba3.Print()

	test1 := testinterface.App{Something: coba3}
	test1.PrintChangenamePrint("ucok")

}
