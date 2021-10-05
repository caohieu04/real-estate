package upcloud

import (
	"context"

	"github.com/caohieu04/real-estate/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
