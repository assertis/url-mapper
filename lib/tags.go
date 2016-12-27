// Copyright 2011 The Go Authors. All rights reserved.

package lib

import (
	"strings"
	"unicode"
)

type TagOptions map[string]string

func (tag TagOptions) Contains(tagName string) bool {
	if _, ok := tag[tagName]; ok {
		return true
	}
	return false
}

// parseTagIntoMap parses a struct tag `valid:required~Some error message,length(2|3)` into map[string]string{"required": "Some error message", "length(2|3)": ""}
func ParseTagsIntoMap(tag string) TagOptions {
	optionsMap := make(TagOptions)
	options := strings.SplitN(tag, ",", -1)
	for _, option := range options {
		validationOptions := strings.Split(option, "=")
		if !isValidTag(validationOptions[0]) {
			continue
		}

		if len(validationOptions) == 2 {
			optionsMap[validationOptions[0]] = validationOptions[1]
		} else {
			optionsMap[validationOptions[0]] = ""
		}
	}
	return optionsMap
}

func isValidTag(s string) bool {
	if s == "" {
		return false
	}

	for _, c := range s {
		switch {
		case strings.ContainsRune("!#$%&()*+-./:<=>?@[]^_{|}~ ", c):
		// Backslash and quote chars are reserved, but
		// otherwise any punctuation chars are allowed
		// in a tag name.
		default:
			if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
				return false
			}
		}
	}
	return true
}
