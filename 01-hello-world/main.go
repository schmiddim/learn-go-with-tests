package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {

	prefix := englishHelloPrefix
	//if language == "Spanish" {
	//	prefix = spanishHelloPrefix
	//}
	//
	//if language == "French" {
	//	prefix = frenchHelloPrefix
	//}

	switch language {
	case "French":
		prefix = frenchHelloPrefix

	case "Spanish":
		prefix = spanishHelloPrefix
	}

	if name == "" {
		name = "World"
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
