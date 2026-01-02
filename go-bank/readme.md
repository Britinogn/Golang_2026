Control structure: Which allows us to know which code execute under the structure
And we use the bank as a case study for this project

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

		//if statement

		//switch
		

		if choice == 1{
			fmt.Println("Your balance is:", accountBalance)

		} else if choice == 2 {
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

		}else if choice == 3{
			fmt.Print("Withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0{
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance{
				fmt.Println("Invalid amount. You can't withdraw more than you have")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Money Withdraw! New amount:", accountBalance)
		}else{
			fmt.Println("Goodbye!")
			break
			//return 
		}	
	}