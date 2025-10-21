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
		// Here you can implement your logic to replace math expressions
		// For demonstration, let's just return the received expression
		// result := replaceMathExpressions(expression)
		c.JSON(http.StatusOK, gin.H{"result": expression})
	})

	router.LoadHTMLFiles("web/index.html")

	router.Run("0.0.0.0:8080")
}
