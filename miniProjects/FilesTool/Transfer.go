package main

import (
	"fmt"
	"os"
)

type Transfer struct {
	Destination string
}

func NewTransfer(destination string) Transfer {
	return Transfer{
		Destination: destination,
	}
}

func (t Transfer) Task(chosen []os.FileInfo) error {
	fmt.Println("Transferring...")
	return nil
}
