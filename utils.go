package main

type number interface {
	int | int32 | int64 | uint | uint32 | uint64
}

func Min[T number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
