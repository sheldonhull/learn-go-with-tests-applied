package hello

const (
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
	germanHelloPrefix   = "Hallo, "
	japeneseHelloPrefix = "Konnichiwa, "
)

// func main() {
// 	fmt.Println(Hello("", ""))
// }

// greetingPrefix returns a string containing the greeting prefix for the given language.
func greetingPrefix(language string) string {
	switch language {
	case "Spanish":
		return spanishHelloPrefix
	case "French":
		return frenchHelloPrefix
	case "German":
		return germanHelloPrefix
	case "Japanese":
		return japeneseHelloPrefix
	default:
		return englishHelloPrefix
	}
}

// Hello returns a string.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}
