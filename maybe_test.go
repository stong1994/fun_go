package gofun_test

import (
	"fmt"
	"github.com/stong1994/gofun"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaybe(t *testing.T) {
	plus2 := gofun.NewMaybe(1).Map(func(value int) int {
		return value + 2
	})
	assert.Equal(t, 3, plus2.Value())

	nothing := gofun.NothingMaybe[int]().Map(func(value int) int {
		return value + 2
	})
	assert.True(t, nothing.IsNothing())
	assert.Equal(t, 0, nothing.Value())

	//upper := func(value string) string { return strings.ToUpper(value) }
	//dup := func(value string) string { return strings.Repeat(value, 2) }
	//ComposeMaybe[string](maybe("some error", upper), dup)("hi")
	withdraw := func(account, amount int) gofun.Maybe[int] {
		if account >= amount {
			return gofun.NewMaybe(account - amount)
		}
		return gofun.InvalidMaybe(0)
	}

	withdrawCurry := gofun.NewCurry(withdraw)
	//withdrawCurry := func(account int) func(amount int) gofun.Maybe[int] {
	//	return func(amount int) gofun.Maybe[int] {
	//		return withdraw(amount, account)
	//	}
	//}

	finishTransaction := func(rst gofun.Maybe[int]) string {
		if rst.IsNothing() {
			return ""
		}
		return fmt.Sprintf("balance is %d", rst.Value())
	}
	getBalance := gofun.MapMaybeElse("you are broke", finishTransaction)
	getTwenty := gofun.UnsafeCompose(getBalance, withdrawCurry.Input(200))
	assert.Equal(t, "balance is 180", getTwenty(20))
	assert.Equal(t, "balance is 180", getTwenty(20))
	assert.Equal(t, "you are broke", getTwenty(220))
}
