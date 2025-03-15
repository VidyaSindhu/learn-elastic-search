package models

type FoodItem struct {
	FoodId int `json:"foodId"`
	FoodName string `json:"food_name"`
	FoodDescription string `json:"food_description"`
	Frice float32 `json:"price"`
	RestaurantId int64 `json:"restaurant_id"`
	FoodType string `json:"food_type"`
}