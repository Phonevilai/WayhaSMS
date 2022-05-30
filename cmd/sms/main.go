package main

import (
	"WayhaSMS/api/middleware"
	"WayhaSMS/api/routes"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("dev.env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
}

var (
	version     = "0.1.0"
	buildcommit = "dev"
	buildtime   = time.Now().String()
)

const idleTimeout = 100 * time.Second

func main() {
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
	app.Get("/dashboard", monitor.New())
	api := app.Group(os.Getenv("VERSION_API"))
	routes.HealthRouter(api)
	routes.AuthRouter(api)
	api.Get("/x", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"API Version": version,
			"buildcommit": buildcommit,
			"buildtime":   buildtime,
		})
	})

	//Protect Tokenz
	smsAPI := api.Group(os.Getenv("SENDSMS_PATH"), middleware.Protected())
	routes.SMSRouter(smsAPI)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
