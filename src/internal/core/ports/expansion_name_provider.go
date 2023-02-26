package ports

import "context"

type ExpansionNameProvider interface {
	// Return a single expansion name with an optional card name.
	ProvideExpansionName(context.Context, string) (string, error)
}
