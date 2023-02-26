package ports

import "context"

type CardNameRepository interface {
	// Returns a list of card names matching the given search prefix
	//
	// You can also pass an optional expansion set name
	CardNameSearchPrefix(context.Context, string, string) ([]string, error)

	CardNameExists(context.Context, string, string) (bool, error)
}
