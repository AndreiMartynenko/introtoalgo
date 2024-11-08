package main

type Person struct {
	Id   int
	Name string
}

// won't change the original struct,
func (p Person) UpdateName(name string) {
	p.Name = name
}

// will change the original struct
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

/*
func main() {
	person := Person{1, "Alice"}
	person.SetName("Alice Thwaites")
	//(&person).SetName("Alice Thwaites")
	//fmt.Println("updated person: %#v\n", person)

	//return

	var acc Account = Account{
		Id:   1,
		Name: "Alice",
		Person: Person{
			Id:   2,
			Name: "Alice T.",
		},
	}
	acc.SetName("My dear Alice T.")

	fmt.Printf("%#v\n", acc)
}

*/
