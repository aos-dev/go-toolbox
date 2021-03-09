package zapcontext

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

var globalFactory *factory

type factory struct {
	sync.Mutex

	fn func() *zap.Logger
}

func (f *factory) Get() *zap.Logger {
	f.Lock()
	defer f.Unlock()

	return f.fn()
}

func (f *factory) Set(fn func() *zap.Logger) {
	f.Lock()
	defer f.Unlock()

	f.fn = fn
}

// SetFactoryFunction is used to change the logger factory function.
func SetFactoryFunction(fn func() *zap.Logger) {
	globalFactory.Set(fn)
}

func init() {
	globalFactory = &factory{
		fn: func() *zap.Logger {
			l, err := zap.NewProduction()
			if err != nil {
				panic(fmt.Errorf("create zap logger: %w", err))
			}
			return l
		},
	}
}
