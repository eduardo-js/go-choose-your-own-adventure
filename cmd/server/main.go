package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eduardo-js/go-choose-your-own-adventure/pkg/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "http server port")
	fileName := flag.String("file", "story.json", "JSON file with 'Choose Your Own Adventure Story'")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	story, err := cyoa.ParseStoryJson(f)
	if err != nil {
		panic(err)
	}
	h := cyoa.NewHandler(story)
	fmt.Printf("Server running @ https://localhost:%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
