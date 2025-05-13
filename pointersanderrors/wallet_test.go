package pointersanderrors

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit should update balance", func(t *testing.T) {
		depositValue := Bitcoin(10)
		wallet := Wallet{}
		wallet.Deposit(depositValue)
		balance := wallet.Balance()
		assertBalance(t, depositValue, balance)
	})

	t.Run("withdraw should update balance", func(t *testing.T) {
		withdrawValue := Bitcoin(10)
		wantBalance := Bitcoin(0)
		wallet := Wallet{balance: withdrawValue}
		err := wallet.Withdraw(withdrawValue)
		balance := wallet.Balance()
		assertBalance(t, wantBalance, balance)
		assertErrorIsNil(t, err)
	})

	t.Run("withdraw should fail if balance less than withdrawValue", func(t *testing.T) {
		startBalance := Bitcoin(5)
		withdrawValue := Bitcoin(10)
		wallet := Wallet{balance: startBalance}
		err := wallet.Withdraw(withdrawValue)
		balance := wallet.Balance()
		assertBalance(t, startBalance, balance)
		assertError(t, ErrInsufficientFunds, err)
	})
}

func assertBalance(t testing.TB, want, got Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("want balance %s but got %s", want, got)
	}
}

func assertError(t testing.TB, want, got error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error but got nil")
	}
	if !errors.Is(got, want) {
		t.Errorf("want err message: %q but got %q", want, got.Error())
	}
}

func assertErrorIsNil(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("expected an nil error but got %q", got.Error())
	}
}
