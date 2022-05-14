package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
}

func main()  {
	
}