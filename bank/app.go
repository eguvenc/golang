package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/eguvenc/bank/files"
)

func main() {
	var accountBalance, err = files.GetBalanceFromFile()
	if err != nil {
		fmt.Println("Error: ")
		fmt.Println(err)
		fmt.Println("-------------------------")
		return
	}
	fmt.Println("Welcome to Go Bank !")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	presentOptions()

	for i := 0; i < 2; i++ {

		var choice int
		fmt.Print("Your choice:")
		fmt.Scan(&choice)

		// if choice == 1 { // }

		switch choice {
		case 1:
			fmt.Println("Your balabce is", accountBalance)
		case 2:
			fmt.Print("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated ! New amount:", accountBalance)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated ! New amount:", accountBalance)
			files.WriteBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye !")
			return
			// break
		}

	}

	fmt.Println("Thanks for choosing our bank")
}
