package service

import (
	"regexp"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(text string) bool {
	text = strings.TrimSpace(text)
	re := regexp.MustCompile(`^[ .-]+$`)
	if !re.MatchString(text) {
		return false
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
	if len(text) < 2 {
		return false
	}
	re := regexp.MustCompile(`[A-Za-zА-Яа-я0-9]`)
	return re.MatchString(text)
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
