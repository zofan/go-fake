package fake

import "fmt"

func RandIP4() string {
	ex := []int{10, 127, 169, 172, 192, 224, 240}

	return fmt.Sprintf(`%d.%d.%d.%d`,
		RandIntNot(1, 255, ex),
		RandIntNot(0, 255, ex),
		RandIntNot(0, 255, ex),
		RandIntNot(0, 255, ex))
}
