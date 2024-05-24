package password

import (
	"math/rand"
	"time"
)

const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits       = "0123456789"
	Symbols      = "!@#$%&*"
)

func Generate(pwdlength int) string {
	if len(LowerLetters)+len(UpperLetters)+len(Digits)+len(Symbols) < pwdlength {
		return ""
	}

	password := ""
	low := LowerLetters
	upper := UpperLetters
	dig := Digits
	sym := Symbols
	allchars := ""

	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)

	for i := 0; i < pwdlength-4; i++ {
		if i == 0 {
			// one lower
			password += getChar(randomizer, &low)
			// one upper
			password += getChar(randomizer, &upper)
			// one digit
			password += getChar(randomizer, &dig)
			// one symbol
			password += getChar(randomizer, &sym)
			allchars = low + upper + dig + sym
		} else {
			// fill
			password += getChar(randomizer, &allchars)
		}
	}
	return password
}

func getChar(randomizer *rand.Rand, chars *string) string {
	idx := randomizer.Intn(len(*chars) - 1)
	item := string((*chars)[idx])
	if len(*chars)-1 > idx {
		*chars = string((*chars)[:idx]) + string((*chars)[idx+1:])
	} else {
		*chars = string((*chars)[:idx])
	}
	return item
}
