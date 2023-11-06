package gofun

import "reflect"

type Either interface {
	IsRight() bool
	Value() any
}

type Left struct {
	value any
}

func NewLeft(value any) Left {
	return Left{value: value}
}

func (l Left) Value() any {
	return l.value
}

func (l Left) Map(f func(any) any) Left {
	return l
}

func (l Left) IsRight() bool {
	return false
}

type Right struct {
	value any
}

func NewRight(value any) Right {
	return Right{value: value}
}

func (r Right) Value() any {
	return r.value
}

func (r Right) Map(f func(any) any) Right {
	return NewRight(f(r.value))
}

func (r Right) IsRight() bool {
	return true
}

// EitherOf :
// f is the function that will called when e.IsRight()==false
// g is the function that will called when e.IsRight()==true
func EitherOf(f, g any, e Either) any {
	if e.IsRight() {
		return reflect.ValueOf(g).Call([]reflect.Value{reflect.ValueOf(e.Value())})[0].Interface()
	} else {
		return reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(e.Value())})[0].Interface()
	}
}

func CurryEither(f, g any) Curry {
	return NewCurry(EitherOf).Input(f, g)
}
