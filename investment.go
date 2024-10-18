package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64 = 5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Number Of Years: ")
	outputText("Number Of Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	var formatedFeatureValue = fmt.Sprintf("Future value: %.1f\n", futureValue)
	var formatedRealFeatureValue = fmt.Sprintf("Future value (Adjusted for Inlation): %.1f\n", futureRealValue)

	// fmt.Printf("Future value: %.1f\nFuture value (Adjusted for Inlation): %.1f\n", futureValue, futureRealValue)
	// fmt.Println("Future value (Adjusted for Inlation): ", futureRealValue)
	fmt.Print(formatedFeatureValue, formatedRealFeatureValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue
}
