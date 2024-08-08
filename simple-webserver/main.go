package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/home", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusAccepted, gin.H{
			"message": "you are in hello world",
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	srv.ListenAndServe()
}
