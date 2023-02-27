package ports

import "context"

type CardLanguageRepository interface {
	AvailableCachedCardLanguages(context.Context) ([]string, error)
}
