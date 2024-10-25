package controllers

import (
	"net/http"
	"test-backend-kemenkeu/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetProducts(c *gin.Context) {
	products, _ := models.GetProducts()

	if products == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": products})
	}
}
func GetProductnById(c *gin.Context) {

	id := c.Param("id")

	product, _ := models.GetProductById(id)
	// if the name is blank we can assume nothing is found
	if product.NamaProduk == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": product})
	}
}
func PostProduct(c *gin.Context) {
	currentTimestamp := time.Now().UnixNano() / int64(time.Microsecond)
	uniqueID := uuid.New().ID()
	var json models.Product
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	json.Id = currentTimestamp + int64(uniqueID)

	success, err := models.AddProduct(json)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func PutProduct(c *gin.Context) {

	id := c.Param("id")

	var json models.Product
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	success, err := models.PutProductById(json, id)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	success, err := models.DeleteProduct(id)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
