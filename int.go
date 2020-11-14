package fake

import "math/rand"

func RandIntNot(min, max int, excludes []int) int {
	for {
		n := rand.Intn(max-min) + min

		for _, e := range excludes {
			if e == n {
				continue
			}
		}

		return n
	}
}
