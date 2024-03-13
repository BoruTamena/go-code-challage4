# Go Challenges

## Challenge 1

1. Create a function in Go called `fileOpener` that opens a file, performs some operations, and then closes the file using the `defer` keyword. Demonstrate the usage of this function by opening a file, writing some data to it, and then calling `fileOpener`.

## Challenge 2

2. Write a function in Go called `incrementer` that returns another function. The returned function should increment a variable defined in the `incrementer` function's scope each time it's called. Demonstrate the closure property by calling the returned function multiple times and printing the incremented values.

## Challenge 3

3. Write a Go program that reads JSON data from a file named "data.json". The JSON data represents an array of objects with two fields: "name" (string) and "age" (integer). Parse this JSON data into a slice of structs in Go and print each person's name and age.

## Challenge 4

4. Implement a maker function that generates counters

```go
func makeCounter() func() int {
}
c1 := makeCounter()
c1() // 0
c1() // 1
