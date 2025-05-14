package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickphan01/short-link/internal/services"
)


func InitRouters() *gin.Engine{
	router := gin.Default()

	router.POST("/short-link", services.Shortlink)
	return router
}