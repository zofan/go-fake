package fake

import "math/rand"

func RandBool() bool {
	return rand.Intn(1) == 1
}
