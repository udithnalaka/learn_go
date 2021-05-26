package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(10)

		actual := wallet.Balance()
		expected := 10

		fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		if actual != expected {
			t.Errorf("actual %d expected %d", actual, expected)
		}

	})
}
