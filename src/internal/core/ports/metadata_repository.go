package ports

import "context"

type MetadataRepository interface {
	SetMetadata(context.Context, string, any) error

	GetMetadata(context.Context, string, any) (bool, error)

	DeleteMetadata(context.Context, string) error
}
