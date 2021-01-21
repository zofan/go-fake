package fake

import "math/rand"

func RandIntNot(min, max int, excludes []int) int {
	for {
		n := Rand(min, max)

		for _, e := range excludes {
			if e == n {
				continue
			}
		}

		return n
	}
}

func Rand(min, max int) int {
	return rand.Intn(max-min) + min
}
