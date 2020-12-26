package main

import (
	"fmt"
	"os"
	"strconv"
)

type Rename struct {
	Path      string
	ToDate    bool
	IdealName string
}

func NewRename(path string, toDate bool, idealName string) Rename {
	return Rename{
		Path:      path,
		ToDate:    toDate,
		IdealName: idealName,
	}
}

func (r Rename) Task(chosen []os.FileInfo) error {
	counter := 1
	var new string
	var e error
	for _, file := range chosen {
		if r.ToDate {
			new = r.Path + file.ModTime().String() + strconv.Itoa(counter)
		} else {
			new = r.Path + r.IdealName + "_" + strconv.Itoa(counter)
		}
		counter++
		e = os.Rename(r.Path+file.Name(), new)
		if e != nil {
			fmt.Println(e)
			return e
		}
	}
	return nil
}
