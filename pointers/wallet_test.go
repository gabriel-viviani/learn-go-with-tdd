package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit test", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(2999))
		assertBalance(t, wallet, Bitcoin(2999))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(2999)}
		wallet.Withdraw(Bitcoin(999))
		assertBalance(t, wallet, Bitcoin(2000))
	})

	t.Run("Withdraw more than balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(999))
		assertBalance(t, wallet, Bitcoin(100))

		if err == nil {
			t.Error("Expected error")
		}
	})
}
