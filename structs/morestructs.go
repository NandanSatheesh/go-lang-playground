package main

import "fmt"

type Contact struct {
	mobileNumber string
	email        string
}

type Person struct {
	firstName string
	lastName  string
	Contact
}

func main() {

	contactDeatils := Contact{"Mobile Number Field", "thisisforspam@gmail.com"}
	person := Person{"Nandan", "Satheesh", contactDeatils}
	fmt.Printf("%+v\n", person)
	person.updateFirstName("Alex")
	person.print()

}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}
