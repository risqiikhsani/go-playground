package testinterface

import "fmt"

// interface for cat and dog
type Animal interface {
	Print()
	ChangeName(s string)
}

// other interface advantages
type App struct {
	Something Animal
}

func (app App) PrintChangenamePrint(q string) {
	fmt.Println("this is called from method")
	fmt.Println("before the name change = ")
	app.Something.Print()
	app.Something.ChangeName(q)
	fmt.Println("after the name change = ")
	app.Something.Print()
}

// other interface advantages
func ChangenamePrintChangenamePrint(a Animal, q1 string, q2 string) {
	fmt.Println("this is called from global function")
	a.ChangeName(q1)
	a.Print()
	a.ChangeName(q2)
	a.Print()
}
