package main

import "fmt"

func main() {
	fmt.Println(Hello("Hello"))
}

// Hello returns a string
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
