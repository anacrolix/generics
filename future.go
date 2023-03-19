package g

import "context"

type Future[T any] interface {
	Wait(ctx context.Context) T
}
