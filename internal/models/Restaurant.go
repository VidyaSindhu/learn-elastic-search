package models

type Restaurant struct {
	RestaurantId int `json:"restaurant_id"`
	RestaurantName string `json:"restaurant_name"`
	Cuisine int64 `json:"cuisine"`
	Address string `json:"address"`
	OwnerId string `json:"ownder_id"`
}