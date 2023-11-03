package gofun_test

import (
	"github.com/stong1994/gofun"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestCompose(t *testing.T) {
	type args[T, K, V any] struct {
		a func(T) K
		b func(K) V
	}
	tests := []struct {
		name  string
		args  args[any, any, any]
		input any
		want  any
	}{
		{
			name: "string",
			args: args[any, any, any]{
				a: func(s any) any { return strings.ToUpper(s.(string)) },
				b: func(s any) any { return s.(string) + "!" },
			},
			input: "hi",
			want:  "HI!",
		},
		{
			name: "int",
			args: args[any, any, any]{
				a: func(s any) any { return s.(int) * 3 },
				b: func(s any) any { return s.(int) + 2 },
			},
			input: 10,
			want:  36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gofun.Compose(tt.args.a, tt.args.b)(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compose() = %v, want %v", got, tt.want)
			}
		})
	}

	// "int2string",
	plus10 := func(s int) int {
		return s + 10
	}
	preStr := func(s int) string {
		return "rst is " + strconv.Itoa(s)
	}
	assert.Equal(t, "rst is 20", gofun.Compose(preStr, plus10)(10))
}

func TestComposes(t *testing.T) {
	exclamation := func(s string) string { return s + "!" }
	upper := func(s string) string { return strings.ToUpper(s) }
	shout := gofun.UnsafeCompose(upper, exclamation)
	assert.Equal(t, "HI!", shout("hi").(string))

	reverse := gofun.UnsafeCompose(gofun.ReverseString, upper, exclamation)
	assert.Equal(t, "!IH", reverse("hi").(string))

	plus10 := func(s int) int { return s + 10 }
	preRst := func(s int) string { return "rst is " + strconv.Itoa(s) }
	shoutPlus := gofun.UnsafeCompose(exclamation, preRst, plus10)
	assert.Equal(t, "rst is 20!", shoutPlus(10))
	assert.Equal(t, "rst is 11!", shoutPlus(1))
}
