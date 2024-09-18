package main

import (
	"net/http"

	"github.com/IqbalFebrianwar/api-golang/src/controllers/product"

	"github.com/IqbalFebrianwar/api-golang/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectData()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"StatusCode": 200,
			"message":    "Welcome to the API! Use /api/",
			"status":     "success",
		})
	})

	r.GET("/api/products", product.GetAll)
	r.GET("/api/products/:id", product.GetById)
	r.POST("/api/products", product.Create)
	r.PATCH("/api/products/:id", product.Update)
	r.DELETE("/api/products/:id", product.Delete)

	r.Run(":3000")
}
