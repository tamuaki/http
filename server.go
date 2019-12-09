package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("*.html")

	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.Run(":8080")
}
