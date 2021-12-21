package main

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {

	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount

}

//method
// struct  Method return Value
func (w *Wallet) Balance() (amount Bitcoin) {
	return w.balance
}
