package leetcodesolutions

import "golang.org/x/exp/constraints"

func max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
