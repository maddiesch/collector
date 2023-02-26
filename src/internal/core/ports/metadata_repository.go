package ports

import "context"

type MetadataRepository interface {
	SetMetadata(context.Context, string, any) error
}
