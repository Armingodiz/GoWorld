package main

import (
	"fmt"
	"sync"
)

/*

func myFunc() {
	fmt.Println("Inside my goroutine")
}

func main() {
	fmt.Println("Hello World")
	go myFunc()
	fmt.Println("Finished Execution")
}


it never actually prints out Inside my goroutine.
 This is because the main function actually terminates before the goroutine gets a chance to execute.


*/

func myFunc(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my goroutine")
	waitgroup.Done()
}

func main() {
	fmt.Println("Hello World")

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go myFunc(&waitgroup)
	waitgroup.Wait()

	fmt.Println("Finished Execution")
}

/* output will be :

Hello World
Inside my goroutine
Finished Execution

*/
