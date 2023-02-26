package ports

import "context"

type ExpansionNameRepository interface {
	// Returns a list of expansion names mathing the given prefix.
	//
	// Additionally, you can provide a card name to help narrow the search results.
	ExpansionNameSearchPrefix(context.Context, string, string) ([]string, error)

	// Returns a boolean indication if the given expansion name exists
	//
	// This should be an exact match
	ExpansionNameExists(context.Context, string) (bool, error)

	// Returns unique names of expansions the given card name has
	ExpansionNameForCard(context.Context, string) ([]string, error)
}
