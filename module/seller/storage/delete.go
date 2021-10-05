package storage

import (
	"context"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/module/seller/model"
)

func (s *sqlStore) SoftDeleteData(ctx context.Context, ID int) error {
	db := s.db

	if err := db.Table(model.Seller{}.TableName()).
		Where("id = ?", ID).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
