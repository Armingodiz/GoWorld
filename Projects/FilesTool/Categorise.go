package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Categorise struct {
	Path string
}

func NewCategorise(path string) *Categorise {
	return &Categorise{
		Path: path,
	}
}

func (c *Categorise) Task(chosen []os.FileInfo) error {

	types := setTypes(chosen)
	createFolders(c.Path, types)
	for _, file := range chosen {
		Original_Path := c.Path + file.Name()
		parts := strings.Split(file.Name(), ".")
		ext := parts[len(parts)-1]
		New_Path := c.Path + ext + "/" + file.Name()
		e := os.Rename(Original_Path, New_Path)
		if e != nil {
			return e
		}
	}
	return nil
}

func createFolders(path string, types []string) error {
	for _, t := range types {
		err := os.Mkdir(path+t, 0755)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}
	return nil
}

func setTypes(files []os.FileInfo) []string {
	var types []string
	for _, file := range files {
		parts := strings.Split(file.Name(), ".")
		ext := parts[len(parts)-1]
		if !checkExistence(ext, types) {
			types = append(types, ext)
		}
	}
	return types
}

func checkExistence(typee string, types []string) bool {
	for _, t := range types {
		if t == typee {
			return true
		}
	}
	return false
}
