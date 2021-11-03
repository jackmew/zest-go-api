package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const responseOk = 200

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(responseOk, gin.H{
			"message": "pong",
		})
	})
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Error().Err(err).Msg("Gin Run failure")
	}
}