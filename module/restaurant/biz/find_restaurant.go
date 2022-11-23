package restaurantbiz

import (
	"context"
	"g05-food-delivery/common"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
)

type FindRestaurantStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

type findRestaurantBiz struct {
	store FindRestaurantStore
}

func NewFindRestaurantBiz(store FindRestaurantStore) *findRestaurantBiz {
	return &findRestaurantBiz{store: store}
}

func (biz *findRestaurantBiz) FindRestaurant(context context.Context, id int) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotFindEntity(restaurantmodel.EntityName, err)
	}
	return data, nil
}
