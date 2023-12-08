package generics

func Equal[T comparable](a, b T) bool {
	return a == b
}

type AllowAdd interface {
	int | float32 | float64 | string
}

type CustomMap[K string | int, V string | int] map[K]V

func Add[T AllowAdd](a, b T) T {
	return a + b
}

type Cmper[T comparable] interface {
	Equal(a, b T) bool
}


