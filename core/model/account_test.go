package model

import (
	"testing"
)

func TestWithdrawal(t *testing.T) {
	acc := Account{
		Balance: 500,
	}

	remaining, err := acc.Withdraw(250)
	if err != nil {
		t.Error("expected being allowed to withdraw $2,50 on an account with R$5,00")
	}

	if remaining == 250 && acc.Balance == 250 {
		return
	}

	t.Error("expected $2,50 as budget after withdrawal of $2,50 on an account with R$5,00")

}

func TestWithdrawalOnAccountWithNoBalance(t *testing.T) {
	acc := Account{
		Balance: 500,
	}

	_, err := acc.Withdraw(500)
	if err != nil {
		t.Error("expected being allowed to withdraw $5,00 on an account with R$5,00")
	}

	_, err = acc.Withdraw(500)
	if err == nil {
		t.Error("expected not being allowed to withdraw $5,00 on an account with R$0")
	}

}

func TestWithdrawalWithNegativeRemaining(t *testing.T) {
	acc := Account{
		Balance: 500,
	}

	_, err := acc.Withdraw(600)
	if err == nil {
		t.Error("expected not being allowed to withdraw $6,00 on an account with R$5,00")
	}

}

func TestNegativeAndZeroWithdrawals(t *testing.T) {
	acc := Account{
		Balance: 500,
	}

	_, err := acc.Withdraw(0)
	if err == nil {
		t.Error("expected error on $0.00 withdrawal")
	}

	_, err = acc.Withdraw(-250)
	if err == nil {
		t.Error("expected error on $-2.50 withdrawal")
	}

}

func TestDeposit(t *testing.T) {
	acc := Account{
		Balance: 500,
	}

	actual := acc.Deposit(250)

	if actual == 750 && acc.Balance == 750 {
		return
	}

	t.Error("expected $7,50 as budget after deposit of $2,50 on an account with R$5,00")

}

func TestBudgetVerification(t *testing.T) {

	noMoneyAccount := Account{
		Balance: 0,
	}

	muchMoneyAccount := Account{
		Balance: 500,
	}

	if has := noMoneyAccount.HasBudget(); has {
		t.Error("expected false on calling HasBudget() on an account with no budget")
	}

	if has := muchMoneyAccount.HasBudget(); !has {
		t.Error("expected true on calling HasBudget() on an account with $5,00")
	}
}
