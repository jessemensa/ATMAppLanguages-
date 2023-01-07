package main

import (
	"fmt"
	"strconv"
)

// SCREEN
type Screen struct{}

func (s *Screen) displayMessage(message string) {
	fmt.Println(message)
}

func (s *Screen) displayMessageLine(message string) {
	fmt.Println(message)
}

func (s *Screen) displayDollarAmount(amount float64) {
	fmt.Printf("$%.2f", amount)
}

// KEYPAD
type Keypad struct {
	input string
}

// func (k *Keypad) GetInput() int {
// 	return strconv.Atoi(k.input)
// }

func (k *Keypad) GetInput() int {
	input, err := strconv.Atoi(k.input)
	if err != nil {
		// Handle the error here.
		// For example, you can print the error and return a default value.
		fmt.Println(err)
		return 0
	}
	return input
}

func NewKeypad() *Keypad {
	var input string
	fmt.Scanln(&input)
	return &Keypad{input: input}
}

// DEPOSIT SLOT
type DepositSlot struct{}

func (d *DepositSlot) IsEnvelopeReceived() bool {
	return true
}

// CASH DISPENSER
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

// ACCOUNT STRUCT
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

// BANKDATABASE
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

// TRANSACTION
// WHAT IS THE DIFFERENCE BETWEEN INTERFACE VS STRUCT ?
type Transaction interface {
	Execute()
	GetAccountNumber() int
	GetScreen() *Screen
	GetBankDatabase() *BankDatabase
}

type transaction struct {
	userAccountNumber int
	atmScreen         *Screen
	atmBankDatabase   *BankDatabase
}

func (t *transaction) Execute() {
	// TODO: Implement the Execute method
}

func (t *transaction) GetAccountNumber() int {
	return t.userAccountNumber
}

func (t *transaction) GetScreen() *Screen {
	return t.atmScreen
}

func (t *transaction) GetBankDatabase() *BankDatabase {
	return t.atmBankDatabase
}

func NewTransaction(userAccountNumber int, atmScreen *Screen, atmBankDatabase *BankDatabase) Transaction {
	return &transaction{
		userAccountNumber: userAccountNumber,
		atmScreen:         atmScreen,
		atmBankDatabase:   atmBankDatabase,
	}
}

// BALANCE INQUIRY
type BalanceInquiry struct {
	Transaction
}

func NewBalanceInquiry(userAccountNumber int, atmScreen *Screen, atmBankDatabase *BankDatabase) Transaction {
	return &BalanceInquiry{
		Transaction: NewTransaction(userAccountNumber, atmScreen, atmBankDatabase),
	}
}

func (b *BalanceInquiry) Execute() {
	bankDatabase := b.GetBankDatabase()
	screen := b.GetScreen()
	availableBalance := bankDatabase.GetAvailableBalance(b.GetAccountNumber())
	totalBalance := bankDatabase.GetTotalBalance(b.GetAccountNumber())

	// display the balance information on the screen
	screen.displayMessageLine("\nBalance Information:")
	screen.displayMessage(" - Available balance: ")
	screen.displayDollarAmount(availableBalance)
	screen.displayMessage("\n - Total balance: ")
	screen.displayDollarAmount(totalBalance)
	screen.displayMessageLine("")
}

// WITHDRAWAL
type Withdrawal struct {
	Transaction
	keypad        *Keypad
	cashDispenser *CashDispenser
	amount        int
	CANCELED      int
}

func NewWithdrawal(userAccountNumber int, atmScreen *Screen, atmBankDatabase *BankDatabase, atmKeypad *Keypad, atmCashDispenser *CashDispenser) Transaction {
	return &Withdrawal{
		Transaction:   NewTransaction(userAccountNumber, atmScreen, atmBankDatabase),
		keypad:        atmKeypad,
		cashDispenser: atmCashDispenser,
		amount:        0,
		CANCELED:      6,
	}
}

func (w *Withdrawal) Execute() {
	cashDispensed := false
	var availableBalance float64
	bankDatabase := w.GetBankDatabase()
	screen := w.GetScreen()

	for !cashDispensed {
		w.amount = w.DisplayMenuOfAmounts()
		if w.amount != w.CANCELED {
			availableBalance = bankDatabase.GetAvailableBalance(w.GetAccountNumber())

			if float64(w.amount) <= availableBalance {
				if w.cashDispenser.IsSufficientCashAvailable(w.amount) {
					bankDatabase.Debit(w.GetAccountNumber(), float64(w.amount))

					w.cashDispenser.DispenseCash(w.amount)
					cashDispensed = true

					screen.displayMessageLine("\nYour cash has been dispensed. Please take your cash now.")
				} else {
					screen.displayMessageLine("\nInsufficient cash available in the ATM. \nPlease choose a smaller amount.")
				}
			} else {
				screen.displayMessageLine("\nInsufficient funds in your account. \nPlease choose a smaller amount.")
			}
		} else {
			screen.displayMessageLine("\nCanceling transaction...")
			return
		}
	}
}

func (w *Withdrawal) DisplayMenuOfAmounts() int {
	userChoice := 0
	amounts := []int{0, 20, 40, 60, 100, 200}

	screen := w.GetScreen()

	for userChoice == 0 {
		screen.displayMessageLine("\nWithdrawal Menu:")
		screen.displayMessageLine("1 - $20")
		screen.displayMessageLine("2 - $40")
		screen.displayMessageLine("3 - $60")
		screen.displayMessageLine("4 - $100")
		screen.displayMessageLine("5 - $200")
		screen.displayMessageLine("6 - Cancel transaction")
		screen.displayMessage("\nChoose a withdrawal amount: ")

		input := w.keypad.GetInput()

		switch input {
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
			userChoice = amounts[input]
			break
		case w.CANCELED:
			userChoice = w.CANCELED
			break
		default:
			screen.displayMessageLine("\nInvalid selection. Try again.")
		}
	}
	return userChoice
}

// DEPOSIT

type Deposit struct {
	Transaction
	keypad      *Keypad
	depositSlot *DepositSlot
	amount      float64
	CANCELED    int
}

func (d *Deposit) Execute() {
	bankDatabase := d.GetBankDatabase()
	screen := d.GetScreen()

	d.amount = d.PromptForDepositAmount()

	if d.amount != float64(d.CANCELED) {
		screen.displayMessage("\nPlease insert a deposit envelope containing ")
		screen.displayDollarAmount(d.amount)
		screen.displayMessageLine(".")

		envelopeReceived := d.depositSlot.IsEnvelopeReceived()

		if envelopeReceived {
			screen.displayMessageLine("\nYour envelope has been received. \nNOTE: The money just deposited will not be available until we verify the amount of any enclosed cash and your checks clear.")
			bankDatabase.Credit(d.GetAccountNumber(), d.amount)
		} else {
			screen.displayMessageLine("\nYou did not insert an envelope, so the ATM has canceled your transaction.")
		}
	} else {
		screen.displayMessageLine("\nCanceling transaction...")
	}
}

func (d *Deposit) PromptForDepositAmount() float64 {
	screen := d.GetScreen()

	screen.displayMessage("\nPlease enter a deposit amount in CENTS (or 0 to cancel): ")
	input := d.keypad.GetInput()

	if input == d.CANCELED {
		return float64(d.CANCELED)
	} else {
		return float64(input) / 100
	}
}
