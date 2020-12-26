package main

import (
  "os"
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
  tool.SetPerformer(NewCategorise(true))
  tool.Performe(nil)
  tool.SetPerformer(NewRename("",true,""))
  tool.Performe(nil)
  tool.SetPerformer(NewTransfer(""))
  tool.Performe(nil)
}
