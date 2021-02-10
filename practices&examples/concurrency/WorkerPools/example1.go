package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
		// This starts up 3 workers, initially blocked because there are no jobs yet.
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	/*
		we use result chan to stop main go routine from finishing before all jobs be done .
		An alternative way is wait groups .
		if we dont use result our output will be nothing !!!
	*/
	for a := 1; a <= numJobs; a++ {
		<-results // ensures that the worker goroutines have finished
	}
}

// output will be :
/*
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5
*/
