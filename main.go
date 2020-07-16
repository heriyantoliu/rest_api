package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// func init() {
// 	initDb()
// }

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	r.GET("/material", getMaterial)
	r.GET("/material/:id", getMaterialDet)
	r.POST("/material", createMaterial)
	r.PUT("/material/:id", updateMaterial)
	r.DELETE("/material/:id", deleteMaterial)
	r.PATCH("/material/:id", releaseMaterial)

	r.Run(":9001")
}
