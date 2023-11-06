package gofun_test

import (
	"github.com/stong1994/gofun"
	"github.com/stretchr/testify/assert"
	"testing"
)

func crossMultiply(a, b, c float64) float64 {
	if a == 0 {
		panic("can not divide by zero")
	}
	return (b * c) / a
}

func TestFunc(t *testing.T) {
	var a, b, c float64 = 100, 420, 10
	curry := gofun.NewCurry(crossMultiply)

	t.Run("partial apply works", func(t *testing.T) {
		partial := curry.Input(a)
		partial = partial.Input(b, c)
		out := partial.Out()

		assert.Equal(t, float64(42), out)
	})
}
