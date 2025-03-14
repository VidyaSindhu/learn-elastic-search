package models

import (
	"encoding/json"
)

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

func Search(queryName string) (string, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"restaurant_name":  queryName,
			},
		},
	}

	jsonQuery, err := json.Marshal(query)
	if err != nil {
		return "", err
	}
	return string(jsonQuery), nil
}
