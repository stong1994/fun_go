package gofun

import "reflect"

// Compose : Compose two functions (a and b), and then execute as a(b()),
// which means they will be executed in the reverse order of the passed arguments.
func Compose[T, K, V any](a func(K) V, b func(T) K) func(T) V {
	return func(data T) V {
		return a(b(data))
	}
}

// UnsafeCompose : Compose only support two functions, while UnsafeCompose support multiple functions.
// It's not safe because it does not specify the parameter types.
func UnsafeCompose(functions ...any) func(data any) any {
	return func(data any) any {
		result := data
		for _, fn := range Reverse(functions) {
			result = reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(result)})[0].Interface()
		}
		return result
	}
}
