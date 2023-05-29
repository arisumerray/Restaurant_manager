package router

import (
	"github.com/gin-gonic/gin"
	"order/internal/dish"
	"order/internal/order"
)

var r *gin.Engine

func InitRouter(dishHandler *dish.Handler, orderHandler *order.Handler) {
	r = gin.Default()

	r.Use(CORS())

	r.GET("/get_dish", dishHandler.GetDish)
	r.GET("/get_dishes", dishHandler.GetAll)
	r.Use(dishHandler.Service.CheckRole("manager"))
	r.Use(dishHandler.Service.UseId())
	r.POST("/create_dish", dishHandler.CreateDish)
	r.PUT("/update_dish", dishHandler.UpdateDish)
	r.DELETE("/delete_dish", dishHandler.DeleteDish)

	r.POST("/create_order", orderHandler.CreateOrder)
	r.GET("/get_order", orderHandler.GetOrder)
}

func Start(addr string) error {
	return r.Run(addr)
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
