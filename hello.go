package main

import (
	"fmt"
	"io"
	"os"
)

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return getGreeting(lang) + name
}

func getGreeting(language string) (prefix string) {

	switch language {
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}

func main() {
	fmt.Println(Hello("", "english"))
	Greet(os.Stdout, "Henna")
}
