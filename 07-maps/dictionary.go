package main

import "errors"

type Dictionary map[string]string

var ErrorNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {

	def, ok := d[word]

	if !ok {
		return "", ErrorNotFound

	}

	return def, nil
}
