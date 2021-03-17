package main

import (
	"fmt"
	"time"
)

func main() {
	N := 5
	outC := New(func(inC chan interface{}) {
		defer close(inC)
		for i := 0; i < N; i++ {
			inC <- i
		}
	}).
		Pipe(func(in interface{}) (interface{}, error) {
			return multiplyTwo(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return square(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return addQuoute(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return addFoo(in.(string)), nil
		}).
		Merge()

	for result := range outC {
		fmt.Printf("Result: %s\n", result)
	}
}

func multiplyTwo(v int) int {
	time.Sleep(2 * time.Second)
	return v * 2
}

func square(v int) int {
	time.Sleep(2 * time.Second)
	return v * v
}

func addQuoute(v int) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("'%d'", v)
}

func addFoo(v string) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("%s - Foo", v)
}
