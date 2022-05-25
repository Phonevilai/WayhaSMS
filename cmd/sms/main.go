package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
}

var (
	version     = "0.1.0"
	buildcommit = "dev"
	buildtime   = time.Now().String()
)

const idleTimeout = 10 * time.Second

func main()  {
	fmt.Println("Hello LDB Bill Payment")
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)

	}
	defer os.Remove("/tmp/live")
	app := fiber.New(fiber.Config{IdleTimeout: idleTimeout})
	// Middleware
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
	Output:     os.Stdout,
		Format:     `${cyan}[${time}] [RequestID:${locals:requestid} ${blue}${method} ${magenta}${host}${path} ${body}] || ${white}[ ResponseStatus ${status} ${reset}${resBody}]` + "\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))
}