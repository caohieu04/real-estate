package storage

import (
	"context"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/module/seller/model"
)

func (s *sqlStore) Create(ctx context.Context, data *model.SellerCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
