package pointersgo

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(wallet *Wallet, expected BitCoin) {
		t.Helper()
		got := wallet.Balance()
		if got != expected {
			t.Errorf("expected %s, got %s", expected, got)
		}
	}

	assertError := func(got, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertNoError := func(err error) {
		if err != nil {
			t.Error("didn't expect an error but got one")
		}
	}

	t.Run("test deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertBalance(&wallet, 10)
	})

	t.Run("test withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(10)

		assertNoError(err)
		assertBalance(&wallet, 10)
	})

	t.Run("insufficient funds", func(t *testing.T) {
		startingBalance := 20
		wallet := Wallet{balance: BitCoin(startingBalance)}
		err := wallet.Withdraw(BitCoin(100))

		assertError(err, ErrInsufficientFunds)
		assertBalance(&wallet, BitCoin(startingBalance))
	})
}
