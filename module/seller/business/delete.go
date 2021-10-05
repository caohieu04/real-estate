package business

import (
	"context"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/module/seller/model"
)

type DeleteSellerStore interface {
	FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*model.Seller, error)
	SoftDeleteData(ctx context.Context, ID int) error
}

type deleteSellerBusiness struct {
	store DeleteSellerStore
}

func NewDeleteSellerBusiness(store DeleteSellerStore) *deleteSellerBusiness {
	return &deleteSellerBusiness{store: store}
}

func (biz *deleteSellerBusiness) DeleteSeller(ctx context.Context, ID int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": ID})

	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(model.EntityName, nil)
	}

	return nil
}
