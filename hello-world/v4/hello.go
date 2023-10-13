package main

import "fmt"

const engilshHelloPrefix = "Hello, "

func Hello(name string) string {
	return engilshHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
