package server

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// 	"github.com/gofiber/fiber/v2/middleware/logger"
// )

// type Server struct {
// 	App *fiber.App
// }

// // func NewServer(todoHandlers handlers.TodoHandler) *Server {
// // }

// func (s *Server) Initialize() {
// 	app := fiber.New()

// 	app.Use(cors.New())
// 	app.Use(logger.New())
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.JSON(fiber.Map{"status": "ok"})
// 	})

// 	todoGroup := app.Group("/todo")
// 	todoGroup.Get("/todo", s.todoHandlers.GetAll)

// 	return app

// }
