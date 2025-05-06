package randomtool

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const (
	lower 	= "lower"
	upper 	= "upper"
	digits 	= "digits"
	special = "special"
)

var Charsets = map[string]string{
	lower:   "abcdefghijklmnopqrstuvwxyz",
	upper:   "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	digits:  "0123456789",
	special: "!@#$%^&*()-_=+[]{}<>?,.",
}

func RandomString (length int, allowedSets []string) (string, error) {
	var builderOfChar strings.Builder
	for _, charset := range allowedSets {
		if chars, ok := Charsets[charset]; ok {
			builderOfChar.WriteString(chars)
		}
	}
	setOfChar := builderOfChar.String()
	if setOfChar == "" {
		return "", errors.New("empty allowed set")
	}
	
	var result strings.Builder

	for i := 0; i < length; i++ {
		idx, err := RandomIndex(len(setOfChar))
		if err != nil {
			return result.String(), err
		}

		result.WriteByte(setOfChar[idx])
	}

	return result.String(), nil
}

func RandomIndex(length int) (int, error) {
	idx, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
	if err != nil {
		return -1, err
	}
	
	return int(idx.Int64()), nil
}