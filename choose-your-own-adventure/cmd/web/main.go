package main

import (
	"flag"
	"fmt"
	"os"

	adventure "github.com/project-with-golang/choose-your-own-adventure"
)

func main() {
	dosya := flag.String("file", "adventure.json", "the json file") //dosya: adventure.json
	flag.Parse()
	fmt.Printf("Using story from %s file\n", *dosya)

	dsy, err := os.Open(*dosya)
	if err != nil {
		fmt.Println(err)
	}

	story, err := adventure.JSONStory(dsy)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", story)
}
