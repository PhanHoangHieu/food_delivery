package ginrestaurant

import (
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	restaurantbiz "g05-food-delivery/module/restaurant/biz"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
	restaurantstorage "g05-food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var id, err = strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}