package types

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type IngredientRequest struct {
	GroceryName string  `json:"groceryname"`
	Quantity    float64 `json:"quantity"`
}

type CreateFoodItemRequest struct {
	Name        string              `json:"name"`
	Price       float64             `json:"price"`
	Ingredients []IngredientRequest `json:"ingredients"`
}

type PlaceOrderRequest struct {
	FoodItemName string `json:"fooditemname"`
	Quantity     uint   `json:"quantity"`
}

type UpdateGroceryRequest struct {
	Quantity float64 `json:"quantity"`
}

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	UserID    string
	UserType  string
	jwt.StandardClaims
}
