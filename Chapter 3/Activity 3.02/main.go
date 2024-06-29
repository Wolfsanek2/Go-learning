// Loan calculator
package main

import (
	"fmt"
)

func main() {
	creditScore := -500
	income := 1000.0
	loanAmount := 1000.0
	loanTerm := 24.0

	if creditScore < 0 {
		fmt.Println("Error")
		return
	}

	goodScore := false
	if creditScore >= 450 {
		goodScore = true
	}

	interestRate := 0.0
	payment := 0.0
	approved := false
	totalCost := 0.0
	if goodScore {
		interestRate = 0.15
		totalCost = loanAmount * (1 + interestRate)
		payment = totalCost / loanTerm
		if payment < income*0.2 {
			approved = true
		}
	} else {
		interestRate = 0.20
		totalCost = loanAmount * (1 + interestRate)
		payment = totalCost / loanTerm
		if payment < income*0.1 {
			approved = true
		}
	}

	fmt.Println(creditScore)
	fmt.Println(income)
	fmt.Println(loanAmount)
	fmt.Println(loanTerm)
	fmt.Println(payment)
	fmt.Println(interestRate)
	fmt.Println(totalCost)
	fmt.Println(approved)
}
