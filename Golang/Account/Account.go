package main

import (
	"fmt"
	"strconv"
)

type Screen struct{}

func (s *Screen) displayMessage(message string) {
	fmt.Println(message)
}

func (s *Screen) displayMessageLine(message string) {
	fmt.Println(message)
}

func (s *Screen) displayDollarAmount(amount float32) {
	fmt.Printf("$%.2f", amount)
}

type Keypad struct {
	input string
}

func (k *Keypad) GetInput() (int, error) {
	return strconv.Atoi(k.input)
}

func NewKeypad() *Keypad {
	var input string
	fmt.Scanln(&input)
	return &Keypad{input: input}
}

type DepositSlot struct{}

func (d *DepositSlot) IsEnvelopeReceived() bool {
	return true
}

type CashDispenser struct {
	count int
}

const INITIAL_COUNT = 500

func (c *CashDispenser) DispenseCash(amount int) {
	billsRequired := amount / 20
	c.count -= billsRequired
}

func (c *CashDispenser) IsSufficientCashAvailable(amount int) bool {
	billsRequired := amount / 20
	if c.count >= billsRequired {
		return true
	} else {
		return false
	}
}

func NewCashDispenser() *CashDispenser {
	return &CashDispenser{count: INITIAL_COUNT}
}

type Account struct {
	AccountNumber    int
	Pin              int
	AvailableBalance float64
	TotalBalance     float64
}

func (a *Account) ValidatePin(userPin int) bool {
	if userPin == a.Pin {
		return true
	} else {
		return false
	}
}

func (a *Account) GetAvailableBalance() float64 {
	return a.AvailableBalance
}

func (a *Account) GetTotalBalance() float64 {
	return a.TotalBalance
}

func (a *Account) Credit(amount float64) {
	a.TotalBalance += amount
}

func (a *Account) Debit(amount float64) {
	a.AvailableBalance -= amount
	a.TotalBalance -= amount
}

func (a *Account) GetAccountNumber() int {
	return a.AccountNumber
}

type BankDatabase struct {
	accounts []*Account
}

func NewBankDatabase() *BankDatabase {
	accounts := []*Account{
		{AccountNumber: 12345, Pin: 54321, AvailableBalance: 1000.0, TotalBalance: 1200.0},
		{AccountNumber: 98765, Pin: 56789, AvailableBalance: 200.0, TotalBalance: 200.0},
	}
	return &BankDatabase{accounts: accounts}
}

func (b *BankDatabase) GetAccount(accountNumber int) *Account {
	for _, currentAccount := range b.accounts {
		if currentAccount.AccountNumber == accountNumber {
			return currentAccount
		}
	}
	return nil
}

func (b *BankDatabase) AuthenticateUser(userAccountNumber, userPin int) bool {
	userAccount := b.GetAccount(userAccountNumber)
	if userAccount != nil {
		return userAccount.ValidatePin(userPin)
	} else {
		return false
	}
}

func (b *BankDatabase) GetAvailableBalance(userAccountNumber int) float64 {
	return b.GetAccount(userAccountNumber).AvailableBalance
}

func (b *BankDatabase) GetTotalBalance(userAccountNumber int) float64 {
	return b.GetAccount(userAccountNumber).TotalBalance
}

func (b *BankDatabase) Credit(userAccountNumber int, amount float64) {
	b.GetAccount(userAccountNumber).Credit(amount)
}

func (b *BankDatabase) Debit(userAccountNumber int, amount float64) {
	b.GetAccount(userAccountNumber).Debit(amount)
}
