package storage

import (
	"context"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/module/seller/model"
)

func (s *sqlStore) UpdateData(ctx context.Context, ID int, data *model.SellerUpdate) error {
	db := s.db

	if err := db.Where("id = ?", ID).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
