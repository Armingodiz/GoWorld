package main

/*
Race conditions are where 2 threads are accessing memory
at the same time, one of which is writing.
 Race conditions occur because of unsynchronized access to shared memory.
*/
import "fmt"

func main() {
	var data int
	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is %v . \n", data)
	}
}

// in line 7 & 10 both go routines are trying to access the data
// so there are 3 possible outputs :
// 1 ) nothing is preneted (line 7 before 10)
// 2 ) the value is 0 (line 10 & 11 before 7)
// 3 ) the value is 1 (line 10 then 7 then 11 )
