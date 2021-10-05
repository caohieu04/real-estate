package business

import (
	"context"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/module/seller/model"
)

type UpdateSellerStore interface {
	FindDataByCondition(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*model.Seller, error)
	UpdateData(ctx context.Context, ID int, data *model.SellerUpdate) error
}

type updateSellerBusiness struct {
	store UpdateSellerStore
}

func NewUpdateSellerBusiness(store UpdateSellerStore) *updateSellerBusiness {
	return &updateSellerBusiness{store: store}
}

func (biz *updateSellerBusiness) UpdateSeller(ctx context.Context, ID int, data *model.SellerUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": ID})

	if err != nil {
		return common.ErrCannotGetEntity(model.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(model.EntityName, err)
	}

	if err := biz.store.UpdateData(ctx, ID, data); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	return nil
}
