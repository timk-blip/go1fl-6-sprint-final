package service

import (
	"regexp"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(text string) bool {
	text = strings.TrimSpace(text)
	for _, r := range text {
		if r != ' ' && r != '.' && r != '-' {
			return false
		}
	}
	words := strings.Split(text, " / ")
	if len(words) <= 0 {
		return false
	}
	for _, word := range words {
		letters := strings.Fields(word)
		for _, letter := range letters {
			if len(letter) > 5 {
				return false
			}
			if !regexp.MustCompile(`^[.-]+$`).MatchString(letter) {
				return false
			}
		}
	}
	return true
}

func isPlainText(text string) bool {
	c := 0
	if !isMorse(text) {
		for _, r := range text {
			if r == ' ' || r == '.' || r == '-' {
				c++
			}
		}
		if c >= len(text)/2 {
			return false
		}
		return true
	}
	return false
}

func IsParseable(text string) string {
	if len(text) <= 0 {
		return ""
	}

	if isMorse(text) {
		strMorse := morse.ToText(text)
		return strMorse
	}
	if isPlainText(text) {
		strTxt := morse.ToMorse(text)
		return strTxt
	}
	return ""
}
