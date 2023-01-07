package main

import (
	"fmt"
	"strconv"
)

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
