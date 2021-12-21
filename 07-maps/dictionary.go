package main

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrorNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {

	def, ok := d[word]

	if !ok {
		return "", ErrorNotFound

	}

	return def, nil
}

func (d Dictionary) Add(word string, value string) error {

	_, err := d.Search(word)

	switch err {
	case nil:
		return ErrWordExists
	case ErrorNotFound:
		d[word] = value
		return nil
	}

	return nil
}
