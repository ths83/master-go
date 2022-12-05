package main

import (
	"fmt"
	"github.com/appliedgocourses/bank"
	"log"
	"os"
)

func usage() {
	fmt.Println(`Usage:

bank create <name>                     Create an account.
bank list                              List all accounts.
bank update <name> <amount>            Deposit or withdraw money.
bank transfer <name> <name> <amount>   Transfer money between two accounts.
bank history <name>                    Show an account's transaction history.
`)
	os.Exit(1)
}

func main() {
	err := bank.Load()
	if err != nil {
		log.Println("Cannot load bank data:", err)
	}
}

func create(name string) *bank.Account {
	return bank.NewAccount(name)
}

func list() string {
	return bank.ListAccounts()
}

func update(name string, amount int) (int, error) {
	if amount == 0 {
		return 0, fmt.Errorf("cannot add 0 to an account")
	}

	account, _ := getAccount(name)

	if amount > 0 {
		return bank.Deposit(account, amount)
	}
	return bank.Withdraw(account, amount)
}

func transfer(from string, to string, amount int) (int, int, error) {
	if amount == 0 {
		return 0, 0, fmt.Errorf("cannot transfer 0 to another account")
	}

	fromAcct, _ := getAccount(from)
	toAcct, _ := getAccount(to)

	return bank.Transfer(fromAcct, toAcct, amount)
}

func history(name string) func() (amt int, bal int, more bool) {
	account, _ := getAccount(name)
	return bank.History(account)
}

func getAccount(name string) (*bank.Account, error) {
	account, err := bank.GetAccount(name)
	if err != nil {
		return nil, fmt.Errorf("bank account does not exist for name", name)
	}
	return account, nil
}
