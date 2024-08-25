package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./cmd/web/templates/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "test.page.gohtml", gin.H{
			"Title": "Main Web Site",
		})
	})

	router.Run(":8080")
}
