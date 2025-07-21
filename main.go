package main

import (
	"log"

	"github.com/11ALX11/calc-arithmetics/cmd"
	"github.com/11ALX11/calc-arithmetics/i18n"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
		return
	}

	if err := i18n.Init(); err != nil {
		log.Fatalf("failed to initialize i18n: %v", err)
	}

	cmd.Execute()
}
