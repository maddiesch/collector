package ports

import "context"

type CollectorNumberRepository interface {
	// Returns a list of available collector numbers for the given card in an expansion.
	CollectorNumberForCard(context.Context, string, string) ([]string, error)
}
