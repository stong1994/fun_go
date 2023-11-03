package gofun

type Container[T any] struct {
	value T
}

func NewContainer[T any](value T) Container[T] {
	return Container[T]{value: value}
}

func (c Container[T]) Value() T {
	return c.value
}

func (c Container[T]) Map(f func(value T) T) Container[T] {
	return NewContainer[T](f(c.value))
}

func MapContainer[T, V any](c Container[T], f func(value T) V) Container[V] {
	return NewContainer(f(c.value))
}

type UnsafeContainer struct {
	value any
}

func NewUnsafeContainer(value any) UnsafeContainer {
	return UnsafeContainer{value: value}
}

func (u UnsafeContainer) Value() any {
	return u.value
}

func (u UnsafeContainer) Map(f func(value any) any) UnsafeContainer {
	return NewUnsafeContainer(f(u.value))
}
