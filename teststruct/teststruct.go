package teststruct

func TestStruct() {
	z := &Kucing{
		Name: "Wi",
		Age:  4,
	}

	z.Print()

	c := Kucing{
		Name: "Bob",
		Age:  2,
	}

	data := []Kucing{
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
}
