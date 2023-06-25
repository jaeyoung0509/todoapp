package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jaeyoung0509/todo/drivers"
	"github.com/jaeyoung0509/todo/internals/core/ports"
	"github.com/jaeyoung0509/todo/internals/core/service"
	"github.com/jaeyoung0509/todo/internals/handlers"
	"github.com/jaeyoung0509/todo/internals/repositories"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func main() {
	viper.AutomaticEnv() // Automatically use environment variables where available

	app := fx.New(
		fx.Provide(
			drivers.NewClient,
			fx.Annotate(repositories.NewTodoAdapter, fx.As(new(ports.TodoRepository))),
			service.NewTodoService,
			handlers.NewTodoHandler,
		),
		fx.Invoke(NewWebServer),
	)
	app.Run()
}

func NewWebServer(lc fx.Lifecycle, todoHandlers *handlers.TodoHandler) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	todoGroup := app.Group("/todo")
	todoGroup.Get("/todo", todoHandlers.GetAll)

	registerHooks(lc, app)
	return app
}

func registerHooks(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := app.Listen(":3000"); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s \n", err)
				}
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)

			// Block until receive signal
			<-c

			// Create a deadline to wait for
			_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Doesn't block if no connections, but will otherwise wait
			// until the timeout deadline
			return app.Shutdown()
		},
	})
}
