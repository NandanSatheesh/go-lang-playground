package main

import "fmt"

type Contact struct {
	mobileNumber string
	email        string
}

type Person struct {
	firstName      string
	lastName       string
	contactDeatils Contact
}

func main() {

	contactDeatils := Contact{"Mobile Number Field", "thisisforspam@gmail.com"}
	person := Person{"Nandan", "Satheesh", contactDeatils}
	fmt.Printf("%+v\n", person)

}
