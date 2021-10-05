package gin

import (
	"net/http"
	"strconv"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/component"
	"github.com/caohieu04/real-estate/module/seller/business"
	"github.com/caohieu04/real-estate/module/seller/model"
	"github.com/caohieu04/real-estate/module/seller/storage"
	"github.com/gin-gonic/gin"
)

func UpdateSeller(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data model.SellerUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := business.NewUpdateSellerBusiness(store)

		if err := biz.UpdateSeller(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessReponse(true))
	}
}
