package service

import (
	"strings"
)

var MorseMap = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..", 'E': ".",
	'F': "..-.", 'G': "--.", 'H': "....", 'I': "..", 'J': ".---",
	'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.", 'O': "---",
	'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
	'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--",
	'Z': "--..",
	'0': "-----", '1': ".----", '2': "..---", '3': "...--", '4': "....-",
	'5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.",
	' ': "/",
}

var CharMap = map[string]string{
	".-": "A", "-...": "B", "-.-.": "C", "-..": "D", ".": "E",
	"..-.": "F", "--.": "G", "....": "H", "..": "I", ".---": "J",
	"-.-": "K", ".-..": "L", "--": "M", "-.": "N", "---": "O",
	".--.": "P", "--.-": "Q", ".-.": "R", "...": "S", "-": "T",
	"..-": "U", "...-": "V", ".--": "W", "-..-": "X", "-.--": "Y",
	"--..": "Z", "-----": "0", ".----": "1", "..---": "2", "...--": "3",
	"....-": "4", ".....": "5", "-....": "6", "--...": "7", "---..": "8",
	"----.": "9", "/": " ",
}

func ToMorse(str string) string {

	Text := strings.ToUpper(str)
	var result []string

	for _, char := range Text {
		if code, ok := MorseMap[char]; ok {
			result = append(result, code)
		}
	}

	return strings.Join(result, " ")
}

func ToChar(morse string) string {
	words := strings.Split(strings.TrimSpace(morse), "   ")
	var result []string

	for _, word := range words {
		chars := strings.Split(word, " ")
		var decodedWord strings.Builder
		for _, char := range chars {
			if val, ok := CharMap[char]; ok {
				decodedWord.WriteString(val)
			}
		}
		result = append(result, decodedWord.String())
	}

	return strings.Join(result, " ")
}

func Detect(s string) string {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return "unknown"
	}

	isMorse := true
	//	hasValidMorseChar := false

	for _, r := range trimmed {
		switch r {
		case '.', '-', ' ', '/':
			continue
		default:
			isMorse = false
		}
	}
	if isMorse {
		return ToChar(s)
	}
	return ToMorse(s)

}
