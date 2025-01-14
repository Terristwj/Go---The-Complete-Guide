package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// ================= VARIABLES =================
	var investmentAmt int
	var years int = 10
	expectedReturnRate := 5.5

	// Single line declaration
	// investmentAmt, years, expectedReturnRate := 1000, 10, 5.5

	// ================= SCAN =================
	fmt.Print("Enter initial amount: ")
	fmt.Scan(&investmentAmt)
	futureValue, futureRealValue := calculateFutureValues1(investmentAmt, years, expectedReturnRate)

	// ================= PRINT =================
	// Print (Way 1)
	// fmt.Printf("Future Value:\t\t %.2f\n", futureValue)
	// fmt.Printf("Future Real Value:\t %.2f", futureRealValue)

	// Print (Way 2)
	// formattedFV := fmt.Sprintf("Future Value:\t\t %.2f\n", futureValue)
	// formattedFRV := fmt.Sprintf("Future Real Value:\t %.2f", futureRealValue)
	// fmt.Print(formattedFV, formattedFRV)

	// Print (Way 3)
	fmt.Printf(
`Future Value: %.2f
Future Real Value: %.2f`,
		futureValue ,futureRealValue,
	)
}

func calculateFutureValues1(investmentAmt int, years int, expectedReturnRate float64) (float64, float64) {
	fv := float64(investmentAmt) * math.Pow((1 + expectedReturnRate/100), float64(years))
	frv := fv / math.Pow(1+inflationRate/100, float64(years))
	return fv, frv
}

func calculateFutureValues2(investmentAmt int, years int, expectedReturnRate float64) (fv float64, frv float64) {
	fv = float64(investmentAmt) * math.Pow((1 + expectedReturnRate/100), float64(years))
	frv = fv / math.Pow(1+inflationRate/100, float64(years))
	// return fv, frv
	return
}


// Run the code:
// $ go run investment_calculator.go
// Run the main code, if module exists:
// $ go run .

// Build module:
// $ go mod init example.com/investment-calculator
// - Generate go.mod file with module name example.com/investment-calculator

// Build executable:
// $ go build
// - Build executable with name investment-calculator.exe

// Run executable:
// $ ./investment-calculator.exe