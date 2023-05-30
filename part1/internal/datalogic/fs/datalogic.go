package fs

import (
	"context"
)

type Datalogic interface {
	GetDataFromJSON(ctx context.Context) error
}
