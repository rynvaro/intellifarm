package main

import (
	"cattleai/db"
	"cattleai/logsys"
	"cattleai/middleware"
	"cattleai/pkg/acharts"
	"cattleai/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// go:embed index.html static
// var index embed.FS

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

	open := r.Group("/api/open")
	open.GET("/indexdata", acharts.IndexDataHandler)

	// init database
	db.Init()
	defer db.Close()

	// stripped, err := fs.Sub(embedFS, "dist/index.html")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// static.Serve("/", static.LocalFile())
	// r.StaticFS("/index/", http.FS(index))

	r.Run(":8090")
}
