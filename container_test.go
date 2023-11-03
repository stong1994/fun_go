package gofun_test

import (
	"github.com/stong1994/gofun"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestContainer(t *testing.T) {
	plus2 := gofun.NewContainer[int](1).Map(func(value int) int {
		return value + 2
	})
	assert.Equal(t, 3, plus2.Value())

	toString := func(val int) string { return strconv.Itoa(val) }
	plus2String := gofun.MapContainer(plus2, toString)
	assert.Equal(t, "3", plus2String.Value())

	shout := gofun.NewContainer[string]("hi").Map(func(value string) string {
		return value + "!"
	}).Map(func(value string) string {
		return strings.ToUpper(value)
	})
	assert.Equal(t, "HI!", shout.Value())

	container := gofun.NewContainer[string]("Hello")
	calcLen := func(val string) int { return len(val) }
	assert.Equal(t, 5, gofun.MapContainer(container, calcLen).Value())
}

func TestUnsafeContainer(t *testing.T) {
	plus2 := gofun.NewUnsafeContainer(1).Map(func(value any) any {
		return value.(int) + 2
	})
	assert.Equal(t, 3, plus2.Value())

	toString := func(val any) any { return strconv.Itoa(val.(int)) }
	plus2String := plus2.Map(toString)
	assert.Equal(t, "3", plus2String.Value())

	shout := gofun.NewUnsafeContainer("hi").Map(func(value any) any {
		return value.(string) + "!"
	}).Map(func(value any) any {
		return strings.ToUpper(value.(string))
	})
	assert.Equal(t, "HI!", shout.Value())

	container := gofun.NewUnsafeContainer("Hello")
	calcLen := func(val any) any { return len(val.(string)) }
	assert.Equal(t, 5, container.Map(calcLen).Value())
}
