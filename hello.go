package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frencHelloPrefix = "Bonjour, "
const french = "fr"
const english = "en"
const spanish = "es"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return getGreetingPrefix(language) + name
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frencHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Printf(Hello("nha", ""))
}
