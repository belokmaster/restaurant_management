package controller

import (
	"github.com/gin-gonic/gin"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ItemsByOrder(id string) (OrderItem []primitive.M, err error) {

}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	}
}
