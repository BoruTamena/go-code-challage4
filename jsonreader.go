package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Person struct {
	Name string `json "name"`
	Age  int    `json "age"`
}

func read_json(path string) ([]Person, error) {

	var data_slice []Person

	data, err := ioutil.ReadFile("data.json")

	if err != nil {

		return data_slice, err
	}

	err = json.Unmarshal(data, &data_slice)

	if err != nil {

		return data_slice, err

	}

	return data_slice, nil

}

func main() {

	data, err := read_json("data.json")

	if err != nil {

		log.Fatal(err)
	}

	for _, dt := range data {
		fmt.Printf("name:%v \n", dt.Name)
		fmt.Printf("age:%v\n", dt.Age)
	}

}
