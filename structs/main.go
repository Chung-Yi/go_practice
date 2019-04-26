package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	FirstName string
	LastName  string
	Contact   contactInfo
}

func main() {
	// alex := person{FirstName: "Alex", LastName: "Anderson"}
	// fmt.Println(alex)
	// var alex person
	// alex.FirstName = "Alex"
	// alex.LastName = "Anderson"
	// fmt.Printf("%+v", alex)
	jim := person{
		FirstName: "Jim",
		LastName:  "Party",
		Contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// fmt.Printf("%+v", jim)
	// jimPointer := &jim
	// jimPointer.updateName("Andy")
	jim.updateName("Andy")

	jim.print()

}

func (pointerToperson *person) updateName(newFirstname string) {
	(*pointerToperson).FirstName = newFirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
