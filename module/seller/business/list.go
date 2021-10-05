package business

import (
	"context"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/module/seller/model"
)

type ListSellerStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]model.Seller, error)
}

type listSellerBusiness struct {
	store ListSellerStore
}

func NewListSellerBusiness(store ListSellerStore) *listSellerBusiness {
	return &listSellerBusiness{store: store}
}
func (biz *listSellerBusiness) ListSeller(ctx context.Context, filter *model.Filter, paging *common.Paging) ([]model.Seller, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return result, err
}
