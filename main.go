package main

import (
	"log"

	"github.com/11ALX11/calc-arithmetics/cmd"
	"github.com/11ALX11/calc-arithmetics/i18n"
)

func main() {
	if err := i18n.Init(); err != nil {
		log.Fatalf(i18n.T("failed to initialize i18n: %v"), err)
	}

	cmd.Execute()
}
