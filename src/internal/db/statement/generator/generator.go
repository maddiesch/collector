package generator

import (
	"fmt"
	"sync/atomic"
)

type Generator interface {
	Generate() (string, []any, error)
}

type ArgumentNameProvider interface {
	Next() string
}

func NewIncrementingArgumentNameProvider() ArgumentNameProvider {
	return &incrementingArgumentNameProvider{}
}

type incrementingArgumentNameProvider struct {
	current uint64
}

func (i *incrementingArgumentNameProvider) Next() string {
	value := atomic.AddUint64(&i.current, 1)
	return fmt.Sprintf("v%d", value)
}
