// Copyright 2011 The Go Authors. All rights reserved.

package mapper

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

func ParseTagsIntoMap(tag string) (string, TagOptions) {
	optionsMap := make(TagOptions)
	name := ""

	options := strings.SplitN(tag, ",", -1)
	for i, option := range options {
		if i == 0 {
			name = option
			continue
		}

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
	return name, optionsMap
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
