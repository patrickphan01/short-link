package services

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickphan01/short-link/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ShortLinkRequest struct {
	Link string `json:"link"`
}

func Shortlink(c *gin.Context) {
	var link ShortLinkRequest
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	conn, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&models.ShortLink{
		Origin:    c.Request.Host,
		RequestIP: c.ClientIP(),
		Path:      link.Link,
		ShortUrl:  _generateRandomString(10),
	})
}

func _generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
