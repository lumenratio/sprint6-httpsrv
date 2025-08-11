package service

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/log6"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ChechStr(str string) (bool, error) {
	// split string to slices
	ss := strings.Split(str, " ")
	// regex for morse and string
	regMorse := `^(\.){1,5}|(-){1,5}$`
	regStr := `^\p{Cyrillic}|[0-9]|(\.){1}`
	// setup regexp
	rM := regexp.MustCompile(regMorse)
	rS := regexp.MustCompile(regStr)

	// Check that a string is morse code or ordinary string
	var result bool // with this variable we will check variable value. If flip will have detected, then we know that a string is malformed
	// check string with regexp
	// morse = always TRUE
	// string = always FALSE
	for i, v := range ss {
		if rM.MatchString(v) {
			if i > 0 && result == false {
				return false, fmt.Errorf("string is not morse code")
			}
			result = true
		} else if rS.MatchString(v) {
			if i > 0 && result == true {
				return false, fmt.Errorf("string is not ordinary string")
			}
			result = false
		}
	}
	if result {
		return true, nil
	} else {
		return false, nil
	}
}

func MorseConvert(strToConv string) (string, error) {
	if strToConv == "" {
		log6.Err.Println("string is empty")
		return "", fmt.Errorf("string is empty")
	}

	result, err := ChechStr(strToConv)
	if err != nil {
		return "", err
	}

	if result {
		log6.Info.Println("Check done: this string is Morse code")
		return morse.ToText(strToConv), nil
	}
	log6.Info.Println("Check done: this string is ordinary string")
	return morse.ToMorse(strToConv), nil
}
