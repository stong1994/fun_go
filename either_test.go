package gofun_test

import (
	"fmt"
	"github.com/stong1994/gofun"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEither(t *testing.T) {
	rainRight := gofun.NewRight("rain").Map(func(a any) any {
		return "b" + a.(string)
	})
	assert.Equal(t, "brain", rainRight.Value())

	rainLeft := gofun.NewLeft("rain").Map(func(a any) any {
		return "b" + a.(string)
	})
	assert.Equal(t, "rain", rainLeft.Value())

	getAge := func(now time.Time, birthdate string) gofun.Either {
		birth, err := time.Parse("2006-01-02", birthdate)
		if err != nil {
			return gofun.NewLeft("Birth date could not be parsed, " + err.Error())
		}
		return gofun.NewRight(now.Year() - birth.Year() + 1)
	}

	t.Run("right", func(t *testing.T) {
		now, _ := time.Parse("2006-01-02", "2023-12-12")
		age := getAge(now, "2005-12-12")
		assert.Equal(t, 19, age.Value())
		assert.True(t, age.IsRight())
	})
	t.Run("left", func(t *testing.T) {
		now, _ := time.Parse("2006-01-02", "2023-12-12")
		age := getAge(now, "invalid age")
		assert.Contains(t, age.Value(), "Birth date could not be parsed, ")
		assert.False(t, age.IsRight())
	})

	//either := gofun.NewCurry(gofun.EitherOf)
	handleErr := func(err string) string {
		return fmt.Sprintf("handle err: %s", err)
	}
	handleAge := func(age int) string {
		return fmt.Sprintf("age is %d", age)
	}

	getAgeCurry := gofun.NewCurry(getAge)
	now, _ := time.Parse("2006-01-02", "2023-12-12")
	age := gofun.UnsafeCompose(gofun.CurryEither(handleErr, handleAge), getAgeCurry.Input(now))
	//age := gofun.UnsafeCompose(either.Input(handleErr, handleAge), getAgeCurry.Input(now))

	t.Run("compose right", func(t *testing.T) {
		assert.Equal(t, "age is 19", age("2005-12-12"))
	})
	t.Run("compose left", func(t *testing.T) {
		assert.Contains(t, age("invalid age"), "handle err: Birth date could not be parsed, ")
	})

}
