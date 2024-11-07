package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id        int
	Name      string
	Developer func(string) string
	Owner     Person
}

type MyLove struct {
	FirstName string
	LastName  string
	Age       int
}

func (ml MyLove) DisplayInfo() {
	fmt.Println("Name: %s %s, Age: %d\n", ml.FirstName, ml.LastName, ml.Age)
}

// func main() {
// 	//full delcaration
// 	var acc Account = Account{
// 		Id:   1,
// 		Name: "Alice",
// 	}
// 	fmt.Println("%#v\n", acc)

// 	//short declaration
// 	acc.Owner = Person{2, "Alice Thwaites", "London"}

// 	fmt.Printf("%#v\n", acc)
// }

func main() {
	myLove := MyLove{
		FirstName: "Alice",
		LastName:  "Thwaites",
		Age:       24,
	}
	myLove.DisplayInfo()
}
