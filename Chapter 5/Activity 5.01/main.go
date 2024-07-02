// calculating the working hours of employees
package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thusday
	Friday
	Saturday
)

func LogHours(dev *Developer, day Weekday, hours int) {
	dev.WorkWeek[day] = hours
}

func HoursWorked(dev *Developer) int {
	var total int
	for _, v := range dev.WorkWeek {
		total += v
	}
	return total
}

func main() {
	dev := Developer{
		Individual: Employee{
			Id:        1,
			FirstName: "Name",
			LastName:  "LastName",
		},
		HourlyRate: 8,
		WorkWeek:   [7]int{0, 1, 2, 3, 4, 5, 6},
	}
	fmt.Println(dev.WorkWeek[Monday])
	fmt.Println(dev.WorkWeek[Tuesday])
	fmt.Println(HoursWorked(&dev))
}
