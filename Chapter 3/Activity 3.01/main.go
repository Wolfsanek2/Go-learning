// Sales tax calculator
package main

import "fmt"

func taxCalculator(cost float64, taxRate float64) float64 {
	return cost * taxRate
}

func main() {
	itemCosts := map[string]float64{
		"Cake":   0.99,
		"Milk":   2.75,
		"Butter": 0.87,
	}
	itemTaxes := map[string]float64{
		"Cake":   0.075,
		"Milk":   0.015,
		"Butter": 0.02,
	}

	var total float64 = 0
	for item, cost := range itemCosts {
		total += taxCalculator(cost, itemTaxes[item])
	}
	fmt.Println(total)
}
