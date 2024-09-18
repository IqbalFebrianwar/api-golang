package product

import (
	"net/http"

	"github.com/IqbalFebrianwar/api-golang/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"StatusCode": 404,
				"message":    "Data tidak ditemukan!",
				"status":     "error",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"StatusCode": 500,
				"message":    "Terjadi kesalahan saat mengambil data!",
				"status":     "error",
				"error":      err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 200,
		"data":       products,
		"message":    "Data ditemukan!",
		"status":     "success",
	})
}

func GetById(c *gin.Context) {
	var products models.Product
	id := c.Param("id")

	if err := models.DB.First(&products, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"StatusCode": 404,
				"message":    "Data tidak ditemukan!",
				"status":     "error",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"StatusCode": 500,
				"message":    "Terjadi kesalahan saat mengambil data!",
				"status":     "error",
				"error":      err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 200,
		"data":       products,
		"message":    "Data ditemukan!",
		"status":     "success",
	})
}

func Create(c *gin.Context) {
	var products models.Product

	if err := c.ShouldBindJSON(&products); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"StatusCode": 400,
			"message":    "Bad Request!",
			"status":     "error",
			"error":      err.Error(),
		})
	}

	models.DB.Create(&products)
	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 200,
		"data":       products,
		"message":    "Data sudah dibuatkan!",
		"status":     "success",
	})
}

func Update(c *gin.Context) {
	var products models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&products); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"StatusCode": 400,
			"message":    "Bad Request!",
			"status":     "error",
			"error":      err.Error(),
		})
	}

	if models.DB.Model(&products).Where("id = ?", id).Updates(&products).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"StatusCode": 400,
			"message":    "Data tidak dapat di Perbarui!",
			"status":     "error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 200,
		"data":       products,
		"message":    "Data sudah di Perbarui!",
		"status":     "success",
	})
}

func Delete(c *gin.Context) {
	var products models.Product
	id := c.Param("id")

	if err := models.DB.First(&products, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"StatusCode": 404,
				"message":    "Produk tidak ditemukan!",
				"status":     "error",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"StatusCode": 500,
				"message":    "Terjadi kesalahan saat mencari produk!",
				"status":     "error",
				"error":      err.Error(),
			})
		}
		return
	}

	if err := models.DB.Delete(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"StatusCode": 500,
			"message":    "Terjadi kesalahan saat menghapus produk!",
			"status":     "error",
			"error":      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 200,
		"message":    "Produk berhasil dihapus!",
		"status":     "success",
	})
}
