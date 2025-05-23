package prototypecopy

import (
	"fmt"
	"testing"
)

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p // copies Name
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func Test(t *testing.T) {
	john := Person{
		"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{},
	}

	// jane := john

	// shallow copy
	// jane.Name = "Jane" // ok

	// jane.Address.StreetAddress = "321 Baker St"

	// fmt.Println(john.Name, john.Address)
	// fmt.Println(jane.Name, jane. Address)

	// what you really want
	jane := john
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}

	jane.Name = "Jane" // ok

	jane.Address.StreetAddress = "321 Baker St"

	fmt.Println(john.Name, john.Address)
	fmt.Println(jane.Name, jane.Address)
}

func CopyMethodTest(t *testing.T) {
	john := Person{
		"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Matt"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
