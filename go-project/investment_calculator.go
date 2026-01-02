package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5


func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue,futureRealValue := calculateFutureValue(investmentAmount,expectedReturnRate,years)
	
	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value(inflation): %.1f\n", futureRealValue)

	// outputs information
	// fmt.Printf("Future value: %v\nFuture value(inflation): %v\n", futureValue, futureRealValue)
	// fmt.Printf("Future value: %.1f\nFuture value(inflation): %.1f", futureValue, futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64)(float64, float64){
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100,years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv ,rfv
}

//variables types and operator