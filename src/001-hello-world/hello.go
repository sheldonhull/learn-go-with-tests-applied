package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("", ""))
}

// greetingPrefix returns a string containingthe greeting prefix for the given language
func greetingPrefix(language string) string {
	switch language {
	case "Spanish":
		return spanishHelloPrefix
	case "French":
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}

// Hello returns a string
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}
