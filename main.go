package main

import (
	"fmt"
)

const salary = 3350.00

func main() {
	calculateExpenses()	
}

func calculateExpenses() {
	expenseGroups := []string{"Essentials", "Education", "Free", "Future"}
	expenseGroupPercentages := []float64{0.55, 0.05, 0.1, 0.3}

	for i, expenseGroup := range expenseGroups {
		fmt.Printf("%s: %.2f\n", expenseGroup, expenseGroupPercentages[i]*salary)
	}
}
