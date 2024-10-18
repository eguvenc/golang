package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64 = 5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Number Of Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	var formatedFeatureValue = fmt.Sprintf("Future value: %.1f\n", futureValue)
	var formatedRealFeatureValue = fmt.Sprintf("Future value (Adjusted for Inlation): %.1f\n", futureRealValue)

	// fmt.Printf("Future value: %.1f\nFuture value (Adjusted for Inlation): %.1f\n", futureValue, futureRealValue)
	// fmt.Println("Future value (Adjusted for Inlation): ", futureRealValue)
	fmt.Print(formatedFeatureValue, formatedRealFeatureValue)
}
