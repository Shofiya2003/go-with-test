package bank

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)
		assertNoError(t, err)
		assertBalance(t, wallet, 10)
	})

	t.Run("insufficient balance", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(100)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

	})

}

func assertError(t testing.TB, err error, want error) {
	t.Helper()

	if err.Error() != want.Error() {
		t.Errorf("got %q want %q", err, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but did not want one")
	}
}

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got %s want %s ", got, want)
	}

}
