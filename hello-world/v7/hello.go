package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const engilshHelloPrefix = "Hello, "
const spanisnhHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := engilshHelloPrefix

	switch language {
	case "French":
		prefix = frenchHelloPrefix

	case "Spanish":
		prefix = spanisnhHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "english"))
}
