package main

func TestWallet(t *testing.T) {

	acc := Wallet{}

	acc.Deposit(Money(10))

	got := acc.Balance()

	want := Money(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
