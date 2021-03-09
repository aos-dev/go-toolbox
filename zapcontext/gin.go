package zapcontext

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const ginContextKey = "zap"

// FromGin gets logger from gin context
// If context is nil, we will panic it.
// If context doesn't have logger, we will create a new one.
func FromGin(c *gin.Context) *zap.Logger {
	if c == nil {
		panic(fmt.Errorf("gin context is nil"))
	}

	// Get logger from gin context with context key.
	v, ok := c.Get(ginContextKey)
	if !ok {
		return globalFactory.Get()
	}

	// Gin context stores value with string, we need to check if the type is match.
	l, ok := v.(*zap.Logger)
	if !ok {
		panic(fmt.Errorf("gin context key matched with wrong type: %v", l))
	}

	return l
}

// WithinGinContext set a zap logger into context
func WithinGin(c *gin.Context, l *zap.Logger) {
	if c == nil {
		panic(fmt.Errorf("gin context is nil"))
	}

	if l == nil {
		panic(fmt.Errorf("logger is nil"))
	}

	c.Set(ginContextKey, l)
}
