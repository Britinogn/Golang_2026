package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data , _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance , _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

//0644 ; this helps with file permission

func main() {
	//we shouldn't store the account balance in numbers , we will be using file
	var accountBalance = getBalanceFromFile()
	
	fmt.Println("Welcome to Go Bank")

	// for key word for looping 
	for {
		
		fmt.Println("what do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice " )
		fmt.Scan(&choice)

		//wantsCheckBalance := choice == 1

		//switch
		switch choice{
			case 1:
				fmt.Println("Your balance is:", accountBalance)
			case 2:
				fmt.Print("Your deposit: ")
				var depositAmount float64
				fmt.Scan(&depositAmount)

				if depositAmount <= 0{
					fmt.Println("Invalid amount. Must be greater than 0")
					//return
					continue
				}

				accountBalance += depositAmount  //accountBalance += accountBalance + depositAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				writeBalanceToFile(accountBalance)

			case 3:
				fmt.Print("Withdrawal amount: ")
				var withdrawAmount float64
				fmt.Scan(&withdrawAmount)

				if withdrawAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0")
					continue
				}

				if withdrawAmount > accountBalance{
					fmt.Println("Invalid amount. You can't withdraw more than you have")
					continue
				}
				accountBalance -= withdrawAmount
				fmt.Println("Money Withdraw! New amount:", accountBalance)
				writeBalanceToFile(accountBalance)
				
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for banking with us")
				return
				//break
		}	
	}
}