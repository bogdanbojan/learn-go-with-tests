package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(20)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))

	})

}
