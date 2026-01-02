package main

import (
	"fmt"
)

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")
	
	ebt , profit, ratio := calculateProfits(revenue, expenses, taxRate);
	formattedEBT := fmt.Sprintf("Ebt value: %.f\n", ebt)
	formattedP := fmt.Sprintf("Profit value: %.f\n", profit)
	formattedR := fmt.Sprintf("Ratio value(Tax-rate): %.f\n", ratio)

	fmt.Print(formattedEBT,formattedP,formattedR)

	
	// fmt.Printf("%.1f\n", ebt)
	// fmt.Printf("%.1f\n",profit)
	// fmt.Printf("%.3f\n",ratio)

}

func calculateProfits(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses 
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt , profit, ratio
} 

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

