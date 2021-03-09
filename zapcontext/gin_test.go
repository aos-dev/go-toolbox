package zapcontext

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFromGin(t *testing.T) {
	t.Run("gin context is nil", func(t *testing.T) {
		assert.Panics(t, func() {
			FromGin(nil)
		})
	})
	t.Run("normal case", func(t *testing.T) {
		c := &gin.Context{}

		assert.NotPanics(t, func() {
			FromGin(c)
		})

		l := FromGin(c)
		assert.NotNil(t, l)
	})
}

func TestWithinGin(t *testing.T) {
	t.Run("gin context is nil", func(t *testing.T) {
		assert.Panics(t, func() {
			WithinGin(nil, globalFactory.Get())
		})
	})
	t.Run("logger is nil", func(t *testing.T) {
		assert.Panics(t, func() {
			WithinGin(&gin.Context{}, nil)
		})
	})
	t.Run("normal case", func(t *testing.T) {
		ctx := &gin.Context{}
		logger := globalFactory.Get()

		assert.NotPanics(t, func() {
			WithinGin(ctx, logger)
		})
	})
}
