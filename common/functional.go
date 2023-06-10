package common

func Map[T, U any](in []T, fn func(T) U) []U {
	out := make([]U, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

func Filter[T any](in []T, fn func(T) bool) []T {
	out := make([]T, 0)
	for _, v := range in {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

func Reduce[T, U any](in []T, fn func(U, T) U, init U) U {
	for _, v := range in {
		init = fn(init, v)
	}
	return init
}

func FlatMap[T, U any](in []T, fn func(T) []U) []U {
	out := make([]U, 0)
	for _, v := range in {
		out = append(out, fn(v)...)
	}
	return out
}

func Some[T any](in []T, fn func(T) bool) bool {
	for _, v := range in {
		if fn(v) {
			return true
		}
	}
	return false
}

func All[T any](in []T, fn func(T) bool) bool {
	for _, v := range in {
		if !fn(v) {
			return false
		}
	}
	return true
}
