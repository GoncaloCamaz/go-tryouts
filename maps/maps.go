package maps

import "errors"

type Dictionary map[string]string

// func Search(dictionary Dictionary, word string) string {
// 	return dictionary[word]
// }

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// we can modify maps without passing an address to it
