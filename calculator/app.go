package main

import (
	"errors"
	"fmt"
)

func main() {
	var taxRate float64 = 25.8
	revenue, err := getUserInput("Revenue : ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	ebt, profit, ratio := caluculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.2f\n", ebt) //  %.2df means = add two decimal place end of the output
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func caluculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be positive number")
	}

	return userInput, nil
}
