package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const ACCOUNT_BALANCE_FILE = "/var/www/golang/tmp/balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(ACCOUNT_BALANCE_FILE)
	if err != nil {
		return 1000, errors.New("failed to find balanced file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	// You can also write it to a file as a whole.
	err := os.WriteFile(ACCOUNT_BALANCE_FILE, []byte(balanceText), 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Error: ")
		fmt.Println(err)
		fmt.Println("-------------------------")
		return
	}
	fmt.Println("Welcome to Go Bank !")

	for i := 0; i < 2; i++ {
		fmt.Println("Welcome to Go Bank !")
		fmt.Println("What do you want to do ?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye !")
			return
			// break
		}

	}

	fmt.Println("Thanks for choosing our bank")
}
