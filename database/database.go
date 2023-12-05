package model

import "context"

type FileStorage interface {
	SaveMetadata(context.Context, *Metadata) (*Metadata, error)
	GetFiles(context.Context) (*Metadata, error)
}
