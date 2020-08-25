package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
)

type urlPath struct {
	path, url string
}

// parsing data
func ymlParser(ymlBytes []byte) (urlPaths []urlPath, err error) {
	err = yaml.Unmarshal(ymlBytes, &urlPaths)
	return
}

func jsonParser(jsonData []byte) (urlPaths []urlPath, err error) {
	err = json.Unmarshal(jsonData, &urlPaths)
	return
}

// turning our data to mapped data
func mapCreator(urlPaths []urlPath) (mappedData map[string]string) {
	mappedData = make(map[string]string)
	for _, value := range urlPaths {
		mappedData[value.path] = value.url
	}
	return
}
func makeHandler(mappedData map[string]string, defaultHandler http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if url, ok := mappedData[request.URL.Path]; ok {
			http.Redirect(writer, request, url, http.StatusMovedPermanently)
		} else {
			defaultHandler.ServeHTTP(writer, request)
		}
	})

}
func makeYmlHandler(ymlData []byte) (handler http.Handler) {
	data, err2 := ymlParser(ymlData)
	checkErr(err2)
	mappedData := mapCreator(data)
	handler = makeHandler(mappedData, makeDefaultMux())
	return
}
func makeJsonHandler(jsonData []byte) (handler http.Handler) {
	data, err2 := jsonParser(jsonData)
	checkErr(err2)
	mappedData := mapCreator(data)
	handler = makeHandler(mappedData, makeDefaultMux())
	return
}
func makeDefaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", wellcome)
	return mux
}
func wellcome(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "wellcom here !!! ")
}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
