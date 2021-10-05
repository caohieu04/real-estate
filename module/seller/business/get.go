package business

import (
	"context"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/module/seller/model"
)

type GetSellerStore interface {
	FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*model.Seller, error)
}

type getSellerBusiness struct {
	store GetSellerStore
}

func NewGetSellerBusiness(store GetSellerStore) *getSellerBusiness {
	return &getSellerBusiness{store: store}
}

func (biz *getSellerBusiness) GetSeller(ctx context.Context, ID int) (*model.Seller, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": ID})

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(model.EntityName, nil)
	}

	return data, err
}
