package gofun

import "reflect"

type Wrapper func(partialArgs ...any) (Wrapper, []reflect.Value)

type Curry struct {
	inputs  []reflect.Value
	fn      reflect.Value
	nextIdx int
}

func NewCurry(fn any) Curry {
	c := Curry{}
	c.fn = reflect.ValueOf(fn)
	if c.fn.Kind() != reflect.Func {
		panic("fn must be a function")
	}
	c.nextIdx = 0
	c.inputs = make([]reflect.Value, c.fn.Type().NumIn(), c.fn.Type().NumIn())
	return c
}

func (c Curry) Out() any {
	if c.nextIdx != c.fn.Type().NumIn() {
		panic("not enough params")
	}
	return c.fn.Call(c.inputs)[0].Interface()
}

func (c Curry) copy() Curry {
	dst := make([]reflect.Value, len(c.inputs))
	copy(dst, c.inputs)
	return Curry{
		inputs:  dst,
		fn:      c.fn,
		nextIdx: c.nextIdx,
	}
}

func (c Curry) Input(partialArgs ...any) Curry {
	curry := c.copy()
	for i := 0; i < len(partialArgs); i++ {
		curry.inputs[curry.nextIdx] = reflect.ValueOf(partialArgs[i])
		curry.nextIdx++
	}
	return curry
}

//func (c Curry) Func(partialArgs ...any) any {
//	curry := c.copy()
//	for i := 0; i < len(partialArgs); i++ {
//		curry.inputs[curry.nextIdx] = reflect.ValueOf(partialArgs[i])
//		curry.nextIdx--
//	}
//	fnType := c.fn.Type()
//	var params, results []reflect.Type
//	for i := 0; i < c.nextIdx; i++ {
//		params = append(params, fnType.In(i))
//	}
//	for i := 0; i < fnType.NumOut(); i++ {
//		results = append(results, fnType.Out(i))
//	}
//
//	return reflect.MakeFunc(
//		reflect.FuncOf(params, results, false),
//		func(args []reflect.Value) (results []reflect.Value) {
//			args = append(args, c.inputs[c.nextIdx+1:]...)
//			return c.fn.Call(args)
//		}).Interface()
//}

//
//func Curry(fn any) (Wrapper, []reflect.Value) {
//	return wrap(reflect.ValueOf(fn))
//}

func wrap(r reflect.Value) (Wrapper, []reflect.Value) {
	numIn := r.Type().NumIn()

	return func(partialArgs ...any) (Wrapper, []reflect.Value) {
		var partial Wrapper

		args := make([]reflect.Value, numIn)
		idx := numIn - 1

		for _, arg := range partialArgs {
			args[idx] = reflect.ValueOf(arg)
			idx--
		}

		if idx == -1 {
			return nil, r.Call(args)
		}

		partial = func(partialArgs ...any) (Wrapper, []reflect.Value) {
			for _, arg := range partialArgs {
				args[idx] = reflect.ValueOf(arg)
				idx++
			}

			if idx == numIn {
				return nil, r.Call(args)
			}

			return partial, nil
		}

		return partial, nil
	}, nil
}
