package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Hello"))
}

// Hello returns a string
func Hello(name string) string {
	return englishHelloPrefix + name
}
