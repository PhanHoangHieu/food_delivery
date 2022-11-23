package ginrestaurant

import (
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	restaurantbiz "g05-food-delivery/module/restaurant/biz"
	restaurantstorage "g05-food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FindRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewFindRestaurantBiz(store)
		data, err := biz.FindRestaurant(c.Request.Context(), int(uid.GetLocalId()))

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		log.Println(data)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
