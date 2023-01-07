package main

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
