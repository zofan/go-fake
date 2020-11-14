package fake

import "math/rand"

func RandString(length int, runes []rune) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func RandAlnumString(length int) string {
	return RandString(length, []rune(`0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`))
}

func RandAlphaString(length int) string {
	return RandString(length, []rune(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`))
}

func RandDigitString(length int) string {
	return RandString(length, []rune(`0123456789`))
}

func RandHexString(length int) string {
	return RandString(length, []rune(`0123456789abcdef`))
}
