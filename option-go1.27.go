//go:build go1.27

package generics

// Returns Some(f(Value)) if me is Some, otherwise None.
func (me Option[V]) Map[U any](f func(V) U) (out Option[U]) {
	if me.Ok {
		out = Some(f(me.Value))
	}
	return
}

// Returns f(Value) if me is Some, otherwise None. Unlike the pre-Go 1.27 method, f may return an
// Option of a different type.
func (me Option[V]) AndThen[U any](f func(V) Option[U]) (out Option[U]) {
	if me.Ok {
		out = f(me.Value)
	}
	return
}
