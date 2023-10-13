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

	if language == spanish {
		return spanisnhHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	return engilshHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", "english"))
}
