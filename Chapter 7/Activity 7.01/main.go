// calculating pay and performance review
package main

import "fmt"

var ratingMap = map[string]int{
	"Excellent":      5,
	"Good":           4,
	"Fair":           3,
	"Poor":           2,
	"Unsatisfactory": 1,
}

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (d Developer) Pay() (string, float64) {
	return fmt.Sprintf("%s %s", d.Individual.FirstName, d.Individual.LastName), d.HourlyRate * d.HoursWorkedInYear
}

func (m Manager) Pay() (string, float64) {
	return fmt.Sprintf("%s %s", m.Individual.FirstName, m.Individual.LastName), m.Salary + (m.Salary * m.CommissionRate)
}

func payDetails(p Payer) {
	fullName, pay := p.Pay()
	fmt.Println(fullName, pay)
}

func (d Developer) calculateReview() float64 {
	var length, sum int
	for _, i := range d.Review {
		switch v := i.(type) {
		case int:
			sum += v
		case string:
			if rating, exists := ratingMap[v]; exists {
				sum += rating
			}
		}
		length++
	}
	return float64(sum) / float64(length)
}

func main() {
	d := Developer{
		Individual: Employee{
			Id:        0,
			FirstName: "Eric",
			LastName:  "Davis",
		},
		HourlyRate:        10,
		HoursWorkedInYear: 480,
		Review: map[string]interface{}{
			"1": 3,
			"2": "Excellent",
		},
	}
	payDetails(d)
	fmt.Println(d.calculateReview())
}
