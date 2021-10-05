package gin

import (
	"net/http"
	"strconv"

	"github.com/caohieu04/real-estate/common"
	"github.com/caohieu04/real-estate/component"
	"github.com/caohieu04/real-estate/module/seller/business"
	"github.com/caohieu04/real-estate/module/seller/storage"
	"github.com/gin-gonic/gin"
)

func GetSeller(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := business.NewGetSellerBusiness(store)

		data, err := biz.GetSeller(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessReponse(data))
	}
}
