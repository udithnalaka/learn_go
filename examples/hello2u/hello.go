package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const sinhala = "Sinhala"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const sinhalaHelloPrefix = "Aayubowan, "
const defaultWorld = "World"

func Hello(name, language string) string {

	if name == "" {
		name = defaultWorld
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	// using If statement
	/* if language == spanish {
		return spanishHelloPrefix + name
	} else if language == french {
		return frenchHelloPrefix + name
	} else if language == sinhala {
		return sinhalaHelloPrefix + name
	} else {
		return englishHelloPrefix + name
	} */

	// using switch statement
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case sinhala:
		prefix = sinhalaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Udith", "Spanish"))
}
