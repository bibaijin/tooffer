package q9

import (
	"fmt"
)

// Min 输出旋转数组的最小元素
func Min(xs []float64) (float64, error) {
	if len(xs) == 0 {
		return 0, fmt.Errorf("len(xs) == 0")
	}

	if len(xs) == 1 {
		return xs[0], nil
	}
}
