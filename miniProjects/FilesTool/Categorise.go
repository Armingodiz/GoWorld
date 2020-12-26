package main

import (
	"fmt"
  "os"
)

type Categorise struct {
	Cut bool
}

func NewCategorise(c bool) *Categorise {
	return &Categorise{
		Cut: c,
	}
}

func (c *Categorise) Task(chosen []os.FileInfo) error {
	fmt.Println("Categorising...")
  return nil
}
