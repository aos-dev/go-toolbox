package zapcontext

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// contextKey is used as key to store logger in context
var contextKey struct{}

// From gets logger from context
// If context is nil, we will panic it.
// If context doesn't have logger, we will create a new one.
func From(ctx context.Context) *zap.Logger {
	if ctx == nil {
		panic(fmt.Errorf("context is nil"))
	}

	l, ok := ctx.Value(contextKey).(*zap.Logger)
	if ok {
		return l
	}

	return globalFactory.Get()
}

// Within set logger into given context and return.
// If context is nil, we will panic it.
// If logger is nil, we will panic it.
func Within(ctx context.Context, l *zap.Logger) context.Context {
	if ctx == nil {
		panic(fmt.Errorf("context is nil"))
	}

	if l == nil {
		panic(fmt.Errorf("logger is nil"))
	}

	return context.WithValue(ctx, contextKey, l)
}
