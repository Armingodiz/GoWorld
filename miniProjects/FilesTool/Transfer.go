package main

import (
	"os"
)

type Transfer struct {
	Source      string
	Destination string
}

func NewTransfer(source, destination string) Transfer {
	return Transfer{
		Source:      source,
		Destination: destination,
	}
}

func (t Transfer) Task(chosen []os.FileInfo) error {
	for _, file := range chosen {
		Original_Path := t.Source + file.Name()
		New_Path := t.Destination + file.Name()
		e := os.Rename(Original_Path, New_Path)
    if e != nil{
      return e
    }
	}
	return nil
}
