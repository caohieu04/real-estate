package gin

import (
	"net/http"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/component"
	"github.com/caohieu04/real-estate/module/seller/business"
	"github.com/caohieu04/real-estate/module/seller/model"
	"github.com/caohieu04/real-estate/module/seller/storage"
	"github.com/gin-gonic/gin"
)

func ListSeller(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter model.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := storage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := business.NewListSellerBusiness(store)

		result, err := biz.ListSeller(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
