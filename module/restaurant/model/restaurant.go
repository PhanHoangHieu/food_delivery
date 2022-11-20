package restaurantmodel

import (
	"errors"
	"g05-food-delivery/common"
	"strings"
)

type RestaurantType string

const TypeNormal RestaurantType = "normal"
const TypePremium RestaurantType = "premium"

type Restaurant struct {
	common.SQLModel
	Name string         `json:"name" gorm:"column:name;"`
	Addr string         `json:"addr" gorm:"column:addr;"`
	Type RestaurantType `json:"type" gorm:"column:type;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
