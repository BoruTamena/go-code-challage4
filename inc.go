package main

import "fmt"

func incrementer() func() int {

	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {

	inc1 := incrementer()

	fmt.Println("Incrementer Start...")

	fmt.Println(inc1())
	fmt.Println(inc1())
	fmt.Println(inc1())

}
