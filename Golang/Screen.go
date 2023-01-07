package main

import "fmt"

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
