package main
import (
	"fmt"
  "os"
)

type Rename struct {
	Pattern   string
	ToDate    bool
	IdealName string
}

func NewRename(pattern string, toDate bool, idealName string) Rename {
	return Rename{
		Pattern:   pattern,
		ToDate:    toDate,
		IdealName: idealName,
	}
}

func (r Rename) Task(chosen []os.FileInfo) error {
	fmt.Println("Renaming ... ")
  return nil
}
