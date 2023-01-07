package main

type DepositSlot struct{}

func (d *DepositSlot) IsEnvelopeReceived() bool {
	return true
}
