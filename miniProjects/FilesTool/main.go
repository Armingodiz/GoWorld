package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
  "strings"
)

type Performer interface {
	Task(chosen []os.FileInfo) error
}

type Tool struct {
	performer Performer
}

func NewTool(performer Performer) *Tool {
	return &Tool{
		performer: performer,
	}
}

func (t *Tool) SetPerformer(performer Performer) {
	t.performer = performer
}

func (t *Tool) Performe(chosen []os.FileInfo) error {
	return t.performer.Task(chosen)
}

func main() {
	tool := NewTool(nil)
	chosen := tool.Find("./sample", "test")
  fmt.Println(chosen[0].Name())
	tool.SetPerformer(NewCategorise(true))
	tool.Performe(nil)
	tool.SetPerformer(NewRename("", true, ""))
	tool.Performe(nil)
	tool.SetPerformer(NewTransfer(""))
	tool.Performe(nil)
}

func (t *Tool) Find(dir, pattern string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
  chosen := []os.FileInfo{}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("###########found files in  dir  : ")
	for _, file := range files {
		fmt.Print(file.Name())
    if strings.Contains(file.Name(),pattern){
      chosen = append(chosen, file)
    }
	}
	fmt.Println("######################################")
	return chosen
}
