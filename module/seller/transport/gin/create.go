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

func CreateSeller(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.SellerCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := business.NewCreateSellerBusiness(store)

		if err := biz.CreateSeller(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessReponse(data))
	}
}
