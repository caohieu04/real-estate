package business

import (
	"context"

	"github.com/caohieu04/real-estate/module/seller/model"
)

type CreateSellerStore interface {
	Create(ctx context.Context, data *model.SellerCreate) error
}

type createSellerBusiness struct {
	store CreateSellerStore
}

func NewCreateSellerBusiness(store CreateSellerStore) *createSellerBusiness {
	return &createSellerBusiness{store: store}
}

func (biz *createSellerBusiness) CreateSeller(ctx context.Context, data *model.SellerCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)

	return err
}
