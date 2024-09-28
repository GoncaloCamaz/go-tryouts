package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

/*
This functions bellow will not work without pointers because
when a function or a method is called, the arguments are copied.
*/
func (w *Wallet) Deposit(amount Bitcoin) {
	// This bellow shows the address of the balance in the Deposit method which is different from the address of the balance in the test
	// because in go, when a function or a method is called, the arguments are copied
	//fmt.Printf("address of balance in Deposit is %p \n", &w.balance)

	// using w.balance or (*w).balance is the same thing
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
