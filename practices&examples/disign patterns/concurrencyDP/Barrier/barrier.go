package main

/*
purpose of barrier pattern == > put up a barrier that nobody passes until we have all the results we need.
*/
/*this a simple app which sends 2 requests and get their answer concurrently , checks the correctness and if
everything was ok == > merge answers and print them
*/

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	urls := []string{"http://httpbin.org/headers", "http://httpbin.org/user-agent"}
	barrier(urls[0], urls[1])
}

// captureBarrierOutput will capture the outputs in stdout and can be used in writing unit test
func captureBarrierOutput(endpoints ...string) string {
	// Create a pipe that allows us to connect
	// an io.Writer interface to an io.Reader interface
	reader, writer, _ := os.Pipe()
	// connecting pipe writer to os.Stdout
	os.Stdout = writer

	out := make(chan string)

	// Copies reader input to a bytes buffer before sending
	// the contents of the buffer through a channel
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		out <- buf.String()
	}()

	barrier(endpoints...)

	// Close the writer to signal the Goroutine that
	// no more input is going to come to it
	writer.Close()
	temp := <-out

	return temp
}

var timeoutMilliseconds = 5000

type barrierResp struct {
	Err  error
	Resp string
}

func barrier(endpoints ...string) {

	requestNumber := len(endpoints)

	in := make(chan barrierResp, requestNumber)
	defer close(in)

	responses := make([]barrierResp, requestNumber)

	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		// Block the execution waiting for data from the in channel
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}

	if !hasError {
		for _, resp := range responses {
			fmt.Println(resp.Resp)
		}
	}

}

// makeRequest performs HTTP GET request to an url and accepts a channel to output barrierResp values
func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}

	// Create the HTTP client and set the timeout
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}

	// Perform the HTTP GET request
	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	// Parse the response to a []byte
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	// Send it through the channel
	res.Resp = string(byt)
	out <- res
}
