package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	/*
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		fmt.Printf("address of balance in test is %p \n", &wallet)

		if got != want {
		 	here by changing the %s to %d, the error message will be more informative because we implemented a Stringer interface
			t.Errorf("got %s want %d", got, want)
		}
	*/

	assertBallance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(100))
		assertBallance(t, wallet, Bitcoin(100))
	})

	t.Run("withdraw with balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		wallet.Withdraw(Bitcoin(10))
		assertBallance(t, wallet, Bitcoin(90))
	})

	t.Run("withdraw without balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds.Error())
		assertBallance(t, wallet, Bitcoin(10))
	})

}
