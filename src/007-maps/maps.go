package mymaps

import (
	"errors"
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("unable to find value in map")
	}

	return definition, nil
}
