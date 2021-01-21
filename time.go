package fake

import (
	"time"
)

func Birthday() time.Time {
	y := Rand(18, 30)
	m := Rand(1, 12)
	d := Rand(1, 31)
	return time.Now().AddDate(-y, -m, -d)
}
