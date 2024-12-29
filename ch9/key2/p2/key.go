package p2

import "context"

type key struct{}

var myKey1 key

func NewContext(ctx context.Context, u any) context.Context {
	return context.WithValue(ctx, myKey1, u)
}

func GetContext(ctx context.Context) any {
	return ctx.Value(myKey1)
}
