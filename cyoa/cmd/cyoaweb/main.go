package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	fileName := flag.String("Story", "story.json", "Pass json file with story")
	port := flag.Int("Port", 3000, "Choose port")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	story, err := cyoa.JsonStory(file)
	if err != nil {
		log.Fatalf("serialisetion error in json file: %v", err)
	}

	h := cyoa.NewHandler(story /*, cyoa.WithStringFn(someFunc), cyoa.WithTemplate(CustomTmpl)*/)
	fmt.Printf("starting on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
