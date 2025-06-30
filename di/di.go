package di

import (
	"fmt"
	"io"
)

func Great(writer io.Writer, name string) error {
	greeting := "Hello, " + name
	_, err := fmt.Fprint(writer, greeting)
	return err
}
