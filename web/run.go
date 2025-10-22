package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/replace-math-expression", func(c *gin.Context) {
		expression := c.PostForm("expression")
		result, err := replaceMathExpressions(expression)
		c.JSON(http.StatusOK, gin.H{"result": result, "error": err})
	})

	router.LoadHTMLFiles("web/index.html")

	router.Run("0.0.0.0:8080")
}
