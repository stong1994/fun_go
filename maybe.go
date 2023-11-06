package gofun

type Maybe[T any] struct {
	valid bool
	value T
}

func NewMaybe[T any](value T) Maybe[T] {
	return Maybe[T]{value: value, valid: true}
}

func NothingMaybe[T any]() Maybe[T] {
	return Maybe[T]{valid: false}
}

func InvalidMaybe[T any](value T) Maybe[T] {
	return Maybe[T]{valid: false, value: value}
}

func (m Maybe[T]) Value() T {
	return m.value
}

func (m Maybe[T]) IsNothing() bool {
	return !m.valid
}

// Map : map m.value with Function f
// Can't set the return value of f to another generic type
func (m Maybe[T]) Map(f func(value T) T) Maybe[T] {
	if m.IsNothing() {
		return m
	}
	return NewMaybe(f(m.value))
}

//
//func MapMaybe[T, V any](m Maybe[T], f func(value  T) V) Maybe[V] {
//	if m.IsNothing() {
//		return m
//	}
//	return NewMaybe(f(m.value))
//}

func MapMaybeElse[T, V any](e V, f func(Maybe[T]) V) func(maybe Maybe[T]) V {
	return func(maybe Maybe[T]) V {
		if maybe.IsNothing() {
			return e
		}
		return f(maybe)
	}
}
