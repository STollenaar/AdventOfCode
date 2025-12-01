package internal

type Point[T comparable] struct {
	Parent *Point[T]
	X, Y   T
	Cost   int
}
