package models

type FoodElasticsearchDocument struct {
	FoodId int `json:"foodId"`
	FoodName string `json:"food_name"`
}

type RestaurantElasticsearchDocument struct {
	RestaurantId int `json:"restaurant_id"`
	RestaurantName string `json:"restaurant_name"`
}

type IndexConfig struct {
	Settings map[string]interface{} `json:"settings"`
	Mappings map[string]interface{} `json:"mappings"`
}