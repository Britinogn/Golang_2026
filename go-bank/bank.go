package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "output/balance.txt"

func main() {
	//we shouldn't store the account balance in numbers , we will be using file
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----")
		//return
		//panic function
		//panic("Can't continue , sorry")
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Your number", randomdata.PhoneNumber(), "has been verified")

	// for key word for looping 
	for {
		
		presentOptions()

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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
				
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for banking with us")
				return
				//break
		}	
	}
}

