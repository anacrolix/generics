package result

import (
	g "github.com/anacrolix/generics"
)

func Ok[T any](ok T) (res g.Result[T]) {
	res.SetOk(ok)
	return
}

// Seems kinda lame since Go can't infer T when returning into place.
func Err[T any](err error) (res g.Result[T]) {
	res.SetErr(err)
	return
}
