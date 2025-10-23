package web

import (
	"net/http"

	"github.com/11ALX11/calc-arithmetics/flags"
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

		flags.Decrypt = c.PostForm("decrypt") == "on"
		flags.Encrypt = c.PostForm("encrypt") == "on"
		flags.KeyPath = c.PostForm("keyPath")
		flags.Unzip = c.PostForm("unzip") == "on"
		flags.UseEvalLib = c.PostForm("useEvalLib") == "on"
		flags.UseFilterRegex = c.PostForm("useFilterRegex") == "on"
		flags.Archive = c.PostForm("archive") == "on"
		flags.DataFileInArchive = c.PostForm("dataFileInArchive")

		result, err := replaceMathExpressions(expression)
		c.JSON(http.StatusOK, gin.H{"result": result, "error": err})
	})

	router.LoadHTMLFiles("web/index.html")

	router.Run("0.0.0.0:8080")
}
