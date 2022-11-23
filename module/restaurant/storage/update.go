package restaurantstorage

import (
	"context"
	"g05-food-delivery/common"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
)

func (s *sqlStore) Update(context context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
