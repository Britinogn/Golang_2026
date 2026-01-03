package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"

)


func GetFloatFromFile(fileName string) (float64, error) {
	data , err := os.ReadFile(fileName)

	if err != nil{
		return 1000 , errors.New("Failed to find file")
	}

	balanceText := string(data)
	value , err := strconv.ParseFloat(balanceText, 64)

	if err != nil{
		return 1000 , errors.New("Failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string){
	err := os.MkdirAll("output", 0755)
	if err != nil{
		fmt.Println("ERROR writing to file:", err)
		return
	}

	balanceText := fmt.Sprint(value)

	err = os.WriteFile(fileName, []byte(balanceText), 0644)
	if err != nil {
		fmt.Println("ERROR writing to file:", err)
		return
	}

	fmt.Println("Successfully wrote result to file")
}

// func writeBalanceToFile(balance float64){
// 	balanceText := fmt.Sprint(balance)
// 	err := os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)

// 	if err != nil {
// 		fmt.Println("ERROR writing to file:", err)
// 	} else {
// 		fmt.Println("Successfully wrote balance to file")
// 	}
// }



//0644 ; this helps with file permission
