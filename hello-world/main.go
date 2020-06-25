package main

import "fmt"

const (
	spanish   = `Spanish`
	french    = `French`
	indonesia = `Indonesia`

	englishHelloPrefix   = `Hello, `
	spanishHelloPrefix   = `Hola, `
	frenchHelloPrefix    = `Bonjour, `
	indonesiaHelloPrefix = `Halo, `
)

func Hello(name, language string) string {
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
		prefix = spanishHelloPrefix
	case indonesia:
		prefix = indonesiaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello(`world`, ``))
}
