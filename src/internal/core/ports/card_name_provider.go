package ports

import "context"

type CardNameProvider interface {
	// Return a single card name given an optional expansion name
	ProvideCardName(context.Context, string) (string, error)
}
