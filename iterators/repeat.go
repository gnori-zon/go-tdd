package iterators

import "strings"

func Repeat(value string, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(value)
	}
	return repeated.String()
}
