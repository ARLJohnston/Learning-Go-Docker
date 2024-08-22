package main

import "fmt"

type Money int

type Account struct {
	balance Money
}

func (m Money) String() string {
	return fmt.Sprintf("£%d", m)
}

func (w *Account) Deposit(amount Money) {
	w.balance += amount
}

func (w *Account) Balance() Money {
	return w.balance
}
