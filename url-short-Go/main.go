package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var pathsFile = flag.String("pathsFile", "paths.yml", "The file containing shortened paths to URL's")

func main() {
	flag.Parse()
	f, err := os.Open(*pathsFile)
	if err != nil {
		log.Fatalf("Could not open file %s", pathsFile)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		log.Fatalf("Could not read file %s", pathsFile)
	}
	ext := filepath.Ext(*pathsFile)
	var handler http.Handler
	if ext == "yml" {
		handler = makeYmlHandler(buf.Bytes())
	} else if ext == "json" {
		handler = makeJsonHandler(buf.Bytes())
	} else {
		fmt.Println("unsupported file type ")
	}
	http.ListenAndServe(":8585", handler)
}
