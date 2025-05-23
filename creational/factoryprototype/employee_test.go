package factoryprototype

import (
	"fmt"
	"testing"
)

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

// functional
func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 60000}
	case Manager:
		return &Employee{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}

func Test(t *testing.T) {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
