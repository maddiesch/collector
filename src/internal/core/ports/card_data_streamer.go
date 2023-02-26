package ports

import "context"

type CardDataStreamer interface {
	StreamCardData(context.Context, string) (<-chan any, error)
}
