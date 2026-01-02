package main

import (
	"errors"
	"fmt"
	"os"
)

/*
GOALS
1) Validate user input
	show error message & exit if invalid if invalid input is provided
	no negative number
	not 0
2) store calculated results into files

*/

// const totalResults = "result.txt"
const totalResults = "output/result.txt"


func main() {
	revenue , err := getUserInput("Revenue: ")
	if err != nil{
		fmt.Println(err)
	}

	expenses , err := getUserInput("Expenses: ")
	if err != nil{
		fmt.Println(err)
	}

	taxRate , err := getUserInput("Tax Rate: ")
	if err != nil{
		fmt.Println(err)
	}

	// var validateUserInput int
	// fmt.Print("Input a number: ")
	// fmt.Scan(&validateUserInput)
	
	ebt , profit, ratio := calculateProfits(revenue, expenses, taxRate);
	formattedEBT := fmt.Sprintf("Ebt value: %.f\n", ebt)
	formattedP := fmt.Sprintf("Profit value: %.f\n", profit)
	formattedR := fmt.Sprintf("Ratio value(Tax-rate): %.f\n", ratio)

	fmt.Print(formattedEBT,formattedP,formattedR)

	writeResultToFile(ebt,profit,ratio)
	// fmt.Printf("%.1f\n", ebt)
	// fmt.Printf("%.1f\n",profit)
	// fmt.Printf("%.3f\n",ratio)

}

func writeResultToFile(ebt, profit, ratio float64) {
	err := os.MkdirAll("output", 0755)
	if err != nil {
		fmt.Println("ERROR creating folder:", err)
		return
	}

	resultText := fmt.Sprintf(
		"EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n",
		ebt,
		profit,
		ratio,
	)

	err = os.WriteFile(totalResults, []byte(resultText), 0644)
	if err != nil {
		fmt.Println("ERROR writing to file:", err)
		return
	}

	fmt.Println("Successfully wrote result to file")
}


func calculateProfits(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses 
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt , profit, ratio
} 

func getUserInput(infoText string) (float64, error) {

	// var validateUserInput int
	// fmt.Print("Input a number: ")
	// fmt.Scan(&validateUserInput)

	
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Invalid input and value must be a positive number")
	}

	// switch userInput{
	// case 1:
	// 	return 0,errors.New("Invalid Input and Must be greater than 0")
	// }

	return userInput, nil
	
	

	
}

