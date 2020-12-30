package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	cyoa "../.."
)

func main() {
	jsonFile := flag.String("story", "gopher.json", "json file gopher story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *jsonFile)

	f, err := os.Open(*jsonFile)
	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", story)
}
