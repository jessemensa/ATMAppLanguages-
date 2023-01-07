package main

type Transaction interface {
	execute()
}

type transaction struct {
	accountNumber int
	screen        *Screen
	bankDatabase  *BankDatabase
}

func NewTransaction(userAccountNumber int, atmScreen *Screen, atmBankDatabase *BankDatabase) Transaction {
	return &transaction{
		accountNumber: userAccountNumber,
		screen:        atmScreen,
		bankDatabase:  atmBankDatabase,
	}
}

func (t *transaction) GetAccountNumber() int {
	return t.accountNumber
}

func (t *transaction) GetScreen() *Screen {
	return t.screen
}

func (t *transaction) GetBankDatabase() *BankDatabase {
	return t.bankDatabase
}
