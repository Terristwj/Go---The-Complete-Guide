package main

import "fmt"

func main() {
	// User input
	revenue := getUserInput("Revenue")
	expenses := getUserInput("Expenses")
	taxRate := getUserInput("Tax Rate")

	// Calculation
	ebt, profit, ratio := calculateData(revenue, expenses, taxRate)

	// Output
	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func getUserInput(text string) (float64){
	var uVal float64
	fmt.Printf("%v: ", text)
	fmt.Scan(&uVal)
	return uVal
}

func calculateData(revenue, expenses, taxRate float64) (float64, float64, float64){
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}