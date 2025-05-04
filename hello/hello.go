package hello

import (
	"fmt"
	"regexp"
)

const world = "World"
const defaultPrefix = "Hello"

var prefixByLanguage = map[string]string{
	"Spanish": "Hola",
	"French":  "Bonjour",
	"English": defaultPrefix,
}

var spacesRegex = regexp.MustCompile(`\s+`)

func Hello(name, language string) string {
	prefix := getPrefix(language)
	if isBlank(name) {
		return buildHelloMessage(prefix, world)
	}
	return buildHelloMessage(prefix, name)
}

func buildHelloMessage(prefix, name string) string {
	return fmt.Sprintf("%s, %s", prefix, name)
}

func getPrefix(language string) string {
	if languagePrefix, ok := prefixByLanguage[language]; ok {
		return languagePrefix
	} else {
		return defaultPrefix
	}
}

func isBlank(name string) bool {
	return name == "" || withoutSpaces(name) == ""
}

func withoutSpaces(value string) string {
	return spacesRegex.ReplaceAllString(value, "")
}
