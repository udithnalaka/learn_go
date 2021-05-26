package wallet

import "fmt"

type Wallet struct {
	balance int
}

// * is the pointer to the wallet instance
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
}

func (w *Wallet) Balance() int {
	return w.balance
}
