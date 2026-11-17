//go:build !go1.27

package generics

// Deprecated: Use option.AndThen, or [Option.AndThen] with Go 1.27, which allows changing the
// type.
func (me Option[V]) AndThen(f func(V) Option[V]) Option[V] {
	if me.Ok {
		return f(me.Value)
	}
	return me
}
