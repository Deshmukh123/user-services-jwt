package router

import (
	"github.com/gin-gonic/gin"
	// "user-service/internal/db"
	"user-service/internal/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// db.ConnectSupabase()

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	return r
}
