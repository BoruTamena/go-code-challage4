package main

import (
	"log"
	"os"
)

func fileOpener(path string) {

	fil, err := os.Open(path) // opening file if exist

	if err != nil {

		// creating a new file if file doesn't exist
		fil, err = os.Create(path)

		if err != nil {

			log.Fatal(err)
		}
	}

	defer fil.Close() // closing file

	content := "writing content to file "

	_, err = fil.WriteString(content) // writing string content file

	if err != nil {
		log.Fatal(err)
		return
	}

}

func main() {

	fileOpener("newfile.txt")

}
