package storage

import (
	"context"

	"github.com/caohieu04/real-estate/common"
)

func (s *sqlStore) ListImages(context context.Context, ids []int, moreKeys ...string) ([]common.Image, error) {
	var result []common.Image

	db := s.db

	db = db.Table(common.Image{}.TableName())

	if err := db.Where("id in (?)", ids).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
