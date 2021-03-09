package zapcontext

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	t.Run("context is nil", func(t *testing.T) {
		assert.Panics(t, func() {
			From(nil)
		})
	})
	t.Run("no logger in context", func(t *testing.T) {
		ctx := context.Background()

		assert.NotPanics(t, func() {
			From(ctx)
		})
	})
}

func TestWithin(t *testing.T) {
	t.Run("context is nil", func(t *testing.T) {
		assert.Panics(t, func() {
			Within(nil, globalFactory.Get())
		})
	})
	t.Run("logger is nil", func(t *testing.T) {
		assert.Panics(t, func() {
			Within(context.Background(), nil)
		})
	})
	t.Run("normal case", func(t *testing.T) {
		ctx := context.Background()
		logger := globalFactory.Get()

		assert.NotPanics(t, func() {
			Within(ctx, logger)
		})
	})
}
