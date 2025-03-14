package models

type FoodItem struct {
	foodId int `json:"foodId"`
	foodName string `json:"food_name"`
	foodDescription string `json:"food_description`
	price float32 `json:"price"`
	restaurantId int64 `json:"restaurant_id"`
	foodType string `json:"food_type"`
}