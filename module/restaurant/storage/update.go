package restaurantstorage

import (
	"context"
	restaurantmodel "g05-food-delivery/module/restaurant/model"
)

func (s *sqlStore) Update(context context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
