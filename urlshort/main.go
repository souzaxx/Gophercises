package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	yamlFile := flag.String("yaml", "", "yaml file in the format of 'path: <path>, url: <url>'")
	jsonFile := flag.String("json", "", "json file in the format of 'path: <path>, url: <url>'")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	if *yamlFile != "" {
		file, err := ioutil.ReadFile(*yamlFile)
		if err != nil {
			panic(err)
		}

		// Build the YAMLHandler using the mapHandler as the
		// fallback
		yamlHandler, err := YAMLHandler(file, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", yamlHandler)
	}
	if *jsonFile != "" {
		file, err := ioutil.ReadFile(*jsonFile)
		if err != nil {
			panic(err)
		}

		// Build the JSONHandler using the mapHandler as the
		// fallback
		jsonHandler, err := JSONHandler(file, mapHandler)
		if err != nil {
			panic(err)
		}
		fmt.Println("Starting the server on :8080")
		http.ListenAndServe(":8080", jsonHandler)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
