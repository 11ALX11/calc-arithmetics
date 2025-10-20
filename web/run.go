package web

import (
	"fmt"

	"github.com/11ALX11/calc-arithmetics/i18n"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println(i18n.T("web not supported yet."))

	router := gin.Default()
	// router.GET("/", )

	// router.Run("0.0.0.0:80")
}
