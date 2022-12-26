package dbresolver

import "context"

type useSourceCtxType struct{}

var useSourceCtxKey = useSourceCtxType(struct{}{})

func UseDBSource(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, useSourceCtxKey, true)
}

func shouldUseSource(ctx context.Context) bool {
	if ctx == nil {
		return false
	}

	v, ok := ctx.Value(useSourceCtxKey).(bool)
	if !ok {
		return false
	}

	return v
}
