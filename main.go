package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	gcs := NewGcs(
		"BUCKET_NAME",
		"OBJECT_PATH",
		"./key.json",
	)
	switch os.Args[1] {
	case "POST":
		f, err := os.Open("sample.txt")
		if err != nil {
			log.Fatal(err)
		}
		err = gcs.PutObject(f)
		if err != nil {
			log.Fatal(err)
		}
	case "GET":
		body, err := gcs.GetObject()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	default:
	}
	log.Println("done")
}
