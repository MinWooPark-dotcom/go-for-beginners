package main

import (
	"log"
	"os"
)

func main() {
	//Create a folder/directory at a full qualified path

	date := "Wed, 01 Dec 2021 18:08:03 +0900"
	mkdirErr := os.Mkdir(date, 0755) // ? 0755
	if mkdirErr != nil {
		log.Fatal(mkdirErr)
	}

	mkdirErr = os.Mkdir(date+"/image", 0755)
	if mkdirErr != nil {
		log.Fatal(mkdirErr)
	}

	mkdirErr = os.Mkdir(date+"/text", 0755)
	if mkdirErr != nil {
		log.Fatal(mkdirErr)
	}
}
