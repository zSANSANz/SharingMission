package routes

import (
	"rumahbelajar-api/controllers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/heroku/x/hmetrics/onload"
	cors "github.com/rs/cors/wrapper/gin"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           int(12 * time.Hour),
	})

	r.Use(corsConfig)

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	v1 := r.Group("api/v1")
	{
		v1.POST("/walikelass", controllers.PostWalikelas)
		v1.GET("/walikelass", controllers.GetWalikelass)
		v1.GET("/walikelass/:id", controllers.GetWalikelas)
		v1.PUT("/walikelass/:id", controllers.UpdateWalikelas)
		v1.DELETE("/walikelass/:id", controllers.DeleteWalikelas)
	}

	return r
}
