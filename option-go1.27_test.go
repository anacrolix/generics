//go:build go1.27

package generics_test

import (
	"strconv"
	"testing"

	g "github.com/anacrolix/generics"
	"github.com/anacrolix/generics/option"
)

func TestOptionMapMethod(t *testing.T) {
	if got := g.Some(42).Map(strconv.Itoa); got != g.Some("42") {
		t.Errorf("Some.Map = %v", got)
	}
	if got := g.None[int]().Map(strconv.Itoa); got.Ok {
		t.Errorf("None.Map = %v", got)
	}
	if got, old := g.Some(1).Map(strconv.Itoa), option.Map(strconv.Itoa, g.Some(1)); got != old {
		t.Errorf("method %v != function %v", got, old)
	}
}

func TestOptionAndThenMethod(t *testing.T) {
	parse := func(s string) g.Option[int] {
		i, err := strconv.Atoi(s)
		return g.OptionFromTuple(i, err == nil)
	}
	if got := g.Some("7").AndThen(parse); got != g.Some(7) {
		t.Errorf("Some valid = %v", got)
	}
	if got := g.Some("x").AndThen(parse); got.Ok {
		t.Errorf("Some invalid = %v", got)
	}
	if got := g.None[string]().AndThen(parse); got.Ok {
		t.Errorf("None = %v", got)
	}
	// Same-type use, as with the pre-Go 1.27 method, still infers.
	double := func(i int) g.Option[int] { return g.Some(i * 2) }
	if got := g.Some(2).AndThen(double); got != g.Some(4) {
		t.Errorf("same-type = %v", got)
	}
}
