package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MENU struct {
	ID primitive.ObjectID `bson:"_id"`
	Oder_date time.Time `json:"order_date" validate:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Order_id string `json:"order_id"`
	Table_id *string `json:"table_id" validate:"required"` 
}