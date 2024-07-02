// calculating the payable amount for employees based on working hours
package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individal  Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func nonLoggedHours() func(int) int {
	time := 0
	f := func(addedTime int) int {
		time += addedTime
		return time
	}
	fmt.Println("nonLoggedHours", time)
	return f
}

func (d *Developer) PayDay() (int, bool) {
	var total int
	for _, v := range d.WorkWeek {
		total += v
	}
	var overtime bool
	if total > 40 {
		overtime = true
		total = (40 + (total-40)*2) * d.HourlyRate
	} else {
		total = total * d.HourlyRate
	}
	return total, overtime
}

func (d *Developer) PayDetails() {
	for i, v := range d.WorkWeek {
		fmt.Printf("%v: %d\n", Weekday(i), v)
	}
	total, overtime := d.PayDay()
	fmt.Printf("Total: %d, overtime: %v", total, overtime)
}

func main() {
	dev := Developer{
		Individal: Employee{
			Id:        0,
			FirstName: "Name",
			LastName:  "LastName",
		},
		HourlyRate: 8,
		WorkWeek:   [7]int{0, 8, 10, 10, 10, 6, 8},
	}
	f := nonLoggedHours()
	f(2)
	f(3)
	f(5)
	dev.PayDetails()
}
