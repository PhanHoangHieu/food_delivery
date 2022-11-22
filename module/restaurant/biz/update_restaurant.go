package restaurantbiz

import (
	"context"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	Update(context context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(context context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := biz.store.Update(context, id, data); err != nil {
		return err
	}
	return nil
}
