package service

import (
	"strings"
"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)







func Detect(s string) string {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return "unknown"
	}

	isMorse := true


	for _, r := range trimmed {
		switch r {
		case '.', '-', ' ', '/':
			continue
		default:
			isMorse = false
		}
	}
	if isMorse {
		return morse.ToText(s)
	}
	return morse.ToMorse(s)

}
