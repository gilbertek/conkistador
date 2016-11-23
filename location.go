package main

import (
	"errors"
	"strings"
)

type Location struct {
	Description string
	Transitions []string
	Events      []string
	Items       []int
}

func (loc *Location) CanGoTo(locName string) bool {
	for _, name := range loc.Transitions {
		if (strings.ToLower(name) == locName) || (strings.ToLower(name[0:3])) == locName[0:3] {
			return true
		}
	}

	return false
}

func FindLocationName(locName string) (string, error) {
	for key, _ := range locationMap {
		if (strings.ToLower(key) == locName) || (strings.ToLower(key[0:3]) == locName[0:3]) {
			return key, nil
		}
	}
	return "", errors.New("Can't find location " + locName)
}