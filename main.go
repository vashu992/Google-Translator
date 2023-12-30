package main

import (
	"fmt"
	"github.com/vashu992/Google-Translator/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("there is something wrong while reading file ")
		return
	}
	router.Routing()
}
