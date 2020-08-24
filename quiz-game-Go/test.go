package main

import (
	"flag"
	"fmt"
)

var (
	env *string
	port *int
)

// Basic flag declarations are available for string, integer, and boolean options.
func init() {
	env = flag.String("env", "development", "a string")
	port = flag.Int("port", 3000, "an int")
}

func main() {

	// Once all flags are declared, call flag.Parse() to execute the command-line parsing.
	flag.Parse()

	// Here we’ll just dump out the parsed options and any trailing positional
	// arguments. Note that we need to dereference the points with e.g. *evn to
	// get the actual option values.
	fmt.Println("env:", *env)
	fmt.Println("port:", *port)

}
