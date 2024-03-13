package main

import (
	"log"
	"os"
)

func fileOpener(path string) {

	fil, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644) // opening file if exist

	if err != nil {

		log.Fatal(err)
	}

	defer fil.Close() // closing file

	content := "writing content to file \n"

	_, err = fil.WriteString(content) // writing string content file

	if err != nil {
		log.Fatal("what happen:", err)
		return
	}

}

func main() {

	fileOpener("newfile.txt")

}
