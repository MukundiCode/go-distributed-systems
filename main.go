package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Account struct {
	name    string
	balance int
}

func (acc Account) canTransfer(amount int) bool {
	return acc.balance >= amount
}

type AccountExistsError struct {
	name string
}

func (err AccountExistsError) Error() string {
	return fmt.Sprintf("exists: %s", err.name)
}

type NoSuchAccountError struct {
	name string
}

func (err NoSuchAccountError) Error() string {
	return fmt.Sprintf("no such account: %s", err.name)
}

type InsufficientFundsError struct {
	name string
}

func (err InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: %s", err.name)
}

type Accounts []Account

func (acs *Accounts) NewAccount(name string) (bool, error) {
	if acs.Exists(name) {
		return false, AccountExistsError{name}
	}
	*acs = append(*acs, Account{name, 0})
	return true, nil
}

func (acs *Accounts) Transfer(from string, to string, amount int) (bool, error) {
	fromAcc, err := acs.Get(from)
	if err != nil {
		return false, err
	}
	toAcc, err := acs.Get(to)
	if err != nil {
		return false, err
	}
	if !fromAcc.canTransfer(amount) {
		return false, InsufficientFundsError{from}
	}

	fromAcc.balance -= amount
	toAcc.balance += amount

	return true, nil
}

func (acs *Accounts) Withdraw(name string, amount int) (bool, error) {
	acc, err := acs.Get(name)
	if err != nil {
		return false, err
	}
	if !acc.canTransfer(amount) {
		return false, InsufficientFundsError{name}
	}

	acc.balance -= amount
	return true, nil
}

func (acs *Accounts) Deposit(name string, amount int) (bool, error) {
	acc, err := acs.Get(name)
	if err != nil {
		return false, err
	}

	acc.balance += amount
	return true, nil
}

func (acs *Accounts) Balance(name string) (int, error) {
	acc, err := acs.Get(name)
	if err != nil {
		return 0, err
	}

	return acc.balance, nil
}

func (acs *Accounts) Exists(name string) bool {
	for _, a := range *acs {
		if a.name == name {
			return true
		}
	}
	return false
}

func (acs *Accounts) Get(name string) (*Account, error) {
	for i := range *acs {
		if (*acs)[i].name == name {
			return &(*acs)[i], nil
		}
	}
	return nil, NoSuchAccountError{name}
}

func (acs *Accounts) Report() {
	for _, a := range *acs {
		fmt.Printf("%s %d\n", a.name, a.balance)
	}
}

// TODO: NewAccount, Deposit, Withdraw and Balance go here.

func amount(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var accounts Accounts
	// TODO: create the store that holds your accounts, keyed by name.
	for sc.Scan() {
		//fmt.Println(accounts)
		f := strings.Fields(sc.Text())
		if len(f) == 0 {
			continue
		}
		switch f[0] {
		case "open":
			// TODO: f[1] is the name. Print "exists: NAME" if it is taken.
			_, err := accounts.NewAccount(f[1])
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		case "deposit":
			// TODO: f[1] is the name, amount(f[2]) the cents.
			_, err := accounts.Deposit(f[1], amount(f[2]))
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		case "withdraw":
			// TODO: f[1] is the name, amount(f[2]) the cents.
			_, err := accounts.Withdraw(f[1], amount(f[2]))
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		case "transfer":
			// TODO: f[1] is the source, f[2] the destination, amount(f[3]) the cents.
			_, err := accounts.Transfer(f[1], f[2], amount(f[3]))
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		case "report":
			// TODO: one "NAME BALANCE" line per account, sorted by name.
			accounts.Report()

		}
	}
	// This line only keeps the imports used while the TODOs are empty (errors is for
	// Withdraw's failure, sort for report). Delete it once your code uses them.
	_, _, _ = fmt.Fprintf, errors.New, sort.Strings
}
