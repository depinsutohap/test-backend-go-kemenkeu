package main

import (
	"fmt"
	controllers "test-backend-kemenkeu/controller"
	middlewares "test-backend-kemenkeu/middleware"
	"test-backend-kemenkeu/models"

	"github.com/gin-gonic/gin"
)

func main() {
	err := models.ConnectDatabase()
	fmt.Println(err)
	r := gin.Default()

	public := r.Group("/api")

	public.POST("/login", controllers.Login)
	public.GET("/products", controllers.GetProducts)
	public.GET("/products/:id", controllers.GetProductnById)

	public.POST("/products", middlewares.JwtAuthMiddleware(), controllers.PostProduct)
	public.PUT("/products/:id", middlewares.JwtAuthMiddleware(), controllers.PutProduct)
	public.DELETE("/products/:id", middlewares.JwtAuthMiddleware(), controllers.DeleteProduct)

	r.Run(":8080")

}
