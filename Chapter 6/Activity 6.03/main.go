// panic on invalid data submission
package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validateRoutingNumber() {
	if d.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}
}

func (d *directDeposit) validateLastName() error {
	if d.lastName == "" {
		fmt.Println(ErrInvalidLastName)
		return ErrInvalidLastName
	}
	return nil
}

func (d *directDeposit) report() {
	fmt.Println(d.lastName)
	fmt.Println(d.firstName)
	fmt.Println(d.bankName)
	fmt.Println(d.routingNumber)
	fmt.Println(d.accountNumber)
}

func main() {
	dep := directDeposit{
		"",
		"first name",
		"bank name",
		1,
		2,
	}
	dep.validateRoutingNumber()
	dep.validateLastName()
	dep.report()
}
