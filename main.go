package main

import (
	"cattleai/db"
	"cattleai/logsys"
	"cattleai/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logsys.Init()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/api/v1")

	router.Register(v1)

	// init database
	db.Init()
	defer db.Close()

	r.Run(":9090")
}
