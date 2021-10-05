package storage

import (
	"context"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/module/seller/model"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]model.Seller, error) {
	var result []model.Seller

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(model.Seller{}.TableName()).Where(conditions).Where("status in (1)")

	if filter.CityID > 0 {
		db = db.Where("city_id = ?", filter.CityID)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
