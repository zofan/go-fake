package fake

import "math/rand"

func String(length int, runes []rune) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func AlnumString(length int) string {
	return String(length, []rune(`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`))
}

func AlphaString(length int) string {
	return String(length, []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`))
}

func DigitString(length int) string {
	return String(length, []rune(`0123456789`))
}

func HexString(length int) string {
	return String(length, []rune(`0123456789abcdef`))
}
