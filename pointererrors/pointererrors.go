package pointererrors

import "fmt"

type Bitcoin int

/*
при инкапсулировании полей структуры,
для работы с ними в методах надо их передовать по укаазтелю
*/
type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() Stringer
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
