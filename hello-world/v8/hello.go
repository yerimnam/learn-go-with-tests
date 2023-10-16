package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	engilshHelloPrefix  = "Hello, "
	spanisnhHelloPrefix = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix

	case spanish:
		prefix = spanisnhHelloPrefix

	default:
		prefix = engilshHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "english"))
}
