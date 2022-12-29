package main

import (
	"cattleai/db"
	"cattleai/logsys"
	"cattleai/middleware"
	"cattleai/router"
	"embed"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed fe/dist
var fe embed.FS

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
	v1.Use(middleware.ParseTenant())
	v1.Use(middleware.SaveOperation())

	router.Register(v1)

	// init database
	db.Init()
	defer db.Close()

	// stripped, err := fs.Sub(fe, "fe/dist")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// r.StaticFS("/", http.FS(stripped))

	r.Run(":8090")
}
