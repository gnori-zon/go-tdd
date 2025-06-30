package maps

import "errors"

type DictionaryErr string

const (
	ErrNotFoundWord     DictionaryErr = "not found word"
	ErrAlreadyExistWord DictionaryErr = "already exist word"
)

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	description, ok := d[word]
	if !ok {
		return "", ErrNotFoundWord
	}
	return description, nil
}

func (d Dictionary) Add(word, description string) error {
	_, err := d.Search(word)
	switch {
	case errors.Is(err, ErrNotFoundWord):
		d[word] = description
		return nil
	case err == nil:
		return ErrAlreadyExistWord
	}
	return err
}

func (d Dictionary) Update(word, description string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}
	d[word] = description
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}
	delete(d, word)
	return nil
}
