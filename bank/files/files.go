package files

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const ACCOUNT_BALANCE_FILE = "/var/www/golang/tmp/balance.txt"

func GetBalanceFromFile() (float64, error) {
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

func WriteBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	// You can also write it to a file as a whole.
	err := os.WriteFile(ACCOUNT_BALANCE_FILE, []byte(balanceText), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
