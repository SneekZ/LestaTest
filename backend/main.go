package main

import (
	"backend/api"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"POST"},
		ExposeHeaders: []string{"Content-Length"},
		AllowHeaders: []string{"x-requested-with"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}

	r.Use(cors.New(corsConfig))

	r.POST("/tfidf_lemm", api.PostUploadFileWithLemm)
	r.POST("/tfidf", api.PostUploadFile)

	r.Run(":8080")
}